package search

import (
	"fmt"
	"strings"
	"time"

	swanclient "github.com/Dataman-Cloud/swan-search/src/util/go-swan"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

const SEARCH_LOAD_DATA_INTERVAL = 1

type Indexer interface {
	Index(prefetchStore *DocumentStorage)
}

type DocumentStorage struct {
	Store map[string]Document
}

func NewDocumentStorage() *DocumentStorage {
	return &DocumentStorage{Store: make(map[string]Document)}
}

func (storage *DocumentStorage) Empty() {
	storage.Store = make(map[string]Document)
}

func (storage *DocumentStorage) Copy() *DocumentStorage {
	copied := NewDocumentStorage()
	for k, v := range storage.Store {
		copied.Store[k] = v
	}
	return copied
}

func (storage *DocumentStorage) Indices() []string {
	indices := make([]string, 0)
	for i, _ := range storage.Store {
		indices = append(indices, i)
	}
	return indices
}

func (storage *DocumentStorage) Set(key string, doc Document) {
	storage.Store[key] = doc
}

func (storage *DocumentStorage) Get(key string) *Document {
	doc, ok := storage.Store[key]
	if !ok {
		return nil
	}
	return &doc
}

func (storage *DocumentStorage) Unset(key string) {
	_, ok := storage.Store[key]
	if ok {
		delete(storage.Store, key)
	}
}

type SearchApi struct {
	Index         []string
	Store         *DocumentStorage
	PrefetchStore *DocumentStorage

	Indexer *SwanIndexer
}

type Document struct {
	AppId   string
	Name    string
	Type    string
	GroupId uint64 `json:"-"`
	Param   map[string]string
}

func (searchApi *SearchApi) ApiRegister(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	searchApi.IndexData()
	for _, client := range searchApi.Indexer.SwanClients {
		go searchApi.ListenSSEService(client)
	}

	searchV1 := router.Group("/search/v1", middlewares...)
	{
		searchV1.GET("/luckysearch", searchApi.Search)
	}
}

func (searchApi *SearchApi) IndexData() {
	log.Infof("Init index data...")
	searchApi.PrefetchStore = NewDocumentStorage()
	searchApi.Store = NewDocumentStorage()
	searchApi.Indexer.Index(searchApi.PrefetchStore)
	searchApi.Index = searchApi.PrefetchStore.Indices()
	searchApi.Store = searchApi.PrefetchStore.Copy()
}
func (searchApi *SearchApi) ListenSSEService(client swanclient.Swan) {
	for {
		defer func() {
			if err := recover(); err != nil {
				searchApi.ListenSSEService(client)
			}
		}()

		log.Infof("start listening events")
		events, err := client.AddEventsListener()
		if err != nil {
			log.Errorf("Failed to register for events, %s", err)
			time.Sleep(time.Duration(5 * time.Second))
			continue
		}
		select {
		case event := <-events:
			log.Infof("Indexer receive event: %s", event)
			searchApi.UpdateIndexer(event)
		}
	}
}

func (searchApi *SearchApi) UpdateIndexer(event *swanclient.Event) {
	switch event.Event {
	case swanclient.EventTypeTaskStateFinished:
		data := event.Data.(*swanclient.TaskInfoEvent)
		searchApi.PrefetchStore.Unset(data.TaskId)
		fmt.Printf("delete task :%s", data.TaskId)
	case swanclient.EventTypeTaskStatePendingOffer:
		data := event.Data.(*swanclient.TaskInfoEvent)
		doc := searchApi.PrefetchStore.Get(data.TaskId)
		if doc == nil {
			taskNum := strings.Split(data.TaskId, "-")[0]
			appId := strings.Split(data.TaskId, "-")[1]
			searchApi.PrefetchStore.Set(data.TaskId, Document{
				ID:   data.TaskId,
				Name: data.TaskId,
				Type: DOCUMENT_TASK,
				Param: map[string]string{
					"AppId":  appId,
					"TaskId": taskNum,
				},
			})
			fmt.Printf("add task:%s", data.TaskId)
		}
	case swanclient.EventTypeAppStateCreating:
		data := event.Data.(*swanclient.AppInfoEvent)
		doc := searchApi.PrefetchStore.Get(data.AppId)
		if doc == nil {
			searchApi.PrefetchStore.Set(data.AppId, Document{
				ID:   data.AppId,
				Name: data.Name,
				Type: DOCUMENT_APP,
				Param: map[string]string{
					"AppId": data.AppId,
				},
			})
			fmt.Printf("add app:%s", data.AppId)
		}
	case swanclient.EventTypeAppStateDeletion:
		data := event.Data.(*swanclient.AppInfoEvent)
		searchApi.PrefetchStore.Unset(data.AppId)
		fmt.Printf("delete app:%s", data.AppId)

	}
	searchApi.Index = searchApi.PrefetchStore.Indices()
	searchApi.Store = searchApi.PrefetchStore.Copy()
}
