package search

import (
	"fmt"
	"strings"

	swanclient "github.com/Dataman-Cloud/swan-search/src/util/go-swan"

	log "github.com/Sirupsen/logrus"
	"github.com/donovanhide/eventsource"
)

const (
	DOCUMENT_APP       = "app"
	DOCUMENT_CONTAINER = "container"
	DOCUMENT_TASK      = "task"
)

type SwanIndexer struct {
	Indexer

	SwanClients []swanclient.Swan
}

func NewSwanIndex(SwanClients []swanclient.Swan) *SwanIndexer {
	return &SwanIndexer{SwanClients: SwanClients}
}

func (indexer *SwanIndexer) Index(store *DocumentStorage) {
	for _, swanClient := range indexer.SwanClients {
		var filter map[string][]string
		if apps, err := swanClient.Applications(filter); err == nil {
			for _, app := range apps {
				store.Set(app.ID+app.Name, Document{
					ID:   app.ID,
					Name: app.Name,
					Type: DOCUMENT_APP,
					Param: map[string]string{
						"AppId": app.ID,
					},
				})
				log.Infof("add app:%s\n", app.ID)
				if appDetail, err := swanClient.GetApplication(app.ID); err == nil {
					for _, task := range appDetail.Tasks {
						taskNum := strings.Split(task.ID, "-")[0]
						store.Set(task.ID, Document{
							ID:   task.ID,
							Name: task.ID,
							Type: DOCUMENT_TASK,
							Param: map[string]string{
								"AppId":  app.ID,
								"TaskId": taskNum,
							},
						})
						log.Infof("add task:%s\n", task.ID)
					}
				} else {
					log.Warnf(fmt.Sprintf("get application [%s] error: %s", app.ID, err))
					continue
				}
			}
		} else {
			log.Warnf("get applications error:", err)
		}
	}
}

func (indexer *SwanIndexer) ListenSSEService() {
	fmt.Println("listening event from swan...")
	for _, client := range indexer.SwanClients {
		events, err := client.AddEventsListener()
		if err != nil {
			log.Fatalf("Failed to register for events, %s", err)
		}

		//		timer := time.After(10 * time.Second)
		done := false

		// Receive events from channel for 60 seconds
		for {
			if done {
				break
			}
			select {
			//	case <-timer:
			//		log.Printf("Exiting the loop")
			//		done = true
			case event := <-events:
				log.Infof("Indexer receive event: %s", event)
				indexer.UpdateIndexer(event)
			}
		}
	}
}

func (indexer *SwanIndexer) UpdateIndexer(event *eventsource.Event) {
	fmt.Printf("start handle event:%s\n", event)
}
