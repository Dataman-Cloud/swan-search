package search

import (
	"fmt"
	"strings"

	swanclient "github.com/Dataman-Cloud/swan-search/src/util/go-swan"
	log "github.com/Sirupsen/logrus"
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
			fmt.Printf("applications:%s\n", apps)
			for _, app := range apps {
				store.Set(app.ID+app.Name, Document{
					ID:   app.ID,
					Name: app.Name,
					Type: DOCUMENT_APP,
					Param: map[string]string{
						"AppId": app.ID,
					},
				})
				if appDetail, err := swanClient.GetApplication(app.ID); err == nil {
					fmt.Printf("appDetail:%s\n", appDetail)
					fmt.Printf("appDetail tasks:%s\n", appDetail.Tasks)
					for _, task := range appDetail.Tasks {
						taskId := strings.Split(task.ID, "-")[0]
						fmt.Println("task id:%s", task.ID)
						store.Set(task.ID, Document{
							ID:   task.ID,
							Name: task.ID,
							Type: DOCUMENT_TASK,
							Param: map[string]string{
								"AppId":  app.ID,
								"TaskId": taskId,
							},
						})
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
