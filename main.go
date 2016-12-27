package main

import (
	"net/http"
	"path"

	"github.com/Dataman-Cloud/swan-search/src/config"
	"github.com/Dataman-Cloud/swan-search/src/search"
	swanclient "github.com/Dataman-Cloud/swan-search/src/util/go-swan"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func LoadConfig() config.Config {
	return config.DefaultConfig()
}

func main() {

	// load config
	searchConfig := LoadConfig()
	searchConfig.Ip = "172.28.128.4"
	searchConfig.Port = "9888"
	router := gin.New()

	searchApi := search.SearchApi{}

	var swanClients []swanclient.Swan

	for _, swanConfig := range searchConfig.Swans {
		swanAddr := path.Join(swanConfig.Scheme, "://", swanConfig.Ip, ":", swanConfig.Port)
		swanClient, err := swanclient.NewClient(swanAddr)
		if err != nil {
			log.Errorf("fails to setup swan client:", err)
			continue
		}
		swanClients = append(swanClients, swanClient)
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
