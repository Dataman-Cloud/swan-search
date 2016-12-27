package main

import (
	"net/http"

	"github.com/Dataman-Cloud/swan-search/src/config"
	"github.com/Dataman-Cloud/swan-search/src/search"
	"github.com/Dataman-Cloud/swan-search/src/util/log"
	"github.com/gin-gonic/gin"
)

func LoadConfig() config.Config {
	return config.DefaultConfig()
}

func main() {

	// load config
	searchConfig := LoadConfig()
	router := gin.New()

	searchApi := search.SearchApi{}
	//searchApi.Indexer = search.NewCraneIndex()
	searchApi.ApiRegister(router)

	searchAddr := searchConfig.Ip + ":" + searchConfig.Port
	server := &http.Server{
		Addr:           searchAddr,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.G(ctx).Fatal("can't start server: ", err)
	}
}
