package main

import (
	"net/http"

	"github.com/Dataman-Cloud/swan-search/src/config"
	"github.com/Dataman-Cloud/swan-search/src/search"
	"github.com/itsjamie/gin-cors"

	swanclient "github.com/Dataman-Cloud/swan/go-swan"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func main() {

	// load config
	searchConfig := config.LoadConfig("./deploy/config.json")
	router := gin.New()
	router.Use(cors.Middleware(cors.Config{
		Origins: "*",
	}))

	searchApi := search.SearchApi{}

	var swanClients []swanclient.Swan

	for _, cluster := range searchConfig.Clusters {
		for cName, cAddrs := range cluster {
			swanClient, err := swanclient.NewClient(cAddrs, cName)
			if err != nil {
				log.Errorf("fails to setup cluster client:%s", cName)
				continue
			}
			swanClients = append(swanClients, swanClient)
		}
	}
	searchApi.Indexer = search.NewSwanIndex(swanClients)
	searchApi.ApiRegister(router)

	searchAddr := searchConfig.Ip + ":" + searchConfig.Port
	server := &http.Server{
		Addr:           searchAddr,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("can't start server: ", err)
	}

}
