package search

import (
	"fmt"
	"strings"

	swanclient "github.com/Dataman-Cloud/swan/go-swan"

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
			for _, app := range apps {
				store.Set(app.ID, Document{
					ID:   app.ID,
					Name: app.Name,
					Type: DOCUMENT_APP,
					Param: map[string]string{
						"AppId":     app.ID,
						"ClusterId": app.ClusterID,
						"RunAs":     app.RunAs,
					},
					ClusterId: app.ClusterID,
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
								"AppId":     app.ID,
								"TaskIndex": taskNum,
								"ClusterId": app.ClusterID,
								"RunAs":     app.RunAs,
							},
							ClusterId: app.ClusterID,
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
