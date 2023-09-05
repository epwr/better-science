package main

import (
	// TODO: should these be renamed without ep on them? Probably?
	"github.com/epwr/better-science/internal/common/eplog" 
	"github.com/epwr/better-science/internal/common/epconfig"
)

func main(){

	cfg, err := config.LoadConfig();
	if err != nil {
		log.Critical("Could not load a ServiceConfig, exiting...")
		return
	}

	log.Info("Host: " + cfg.Host)
}


