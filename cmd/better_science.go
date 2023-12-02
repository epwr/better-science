package main

import (
	"github.com/epwr/better-science/internal/common/log" 
	"github.com/epwr/better-science/internal/common/config"
)

func main(){

	cfg, err := config.LoadConfig();
	if err != nil {
		log.Critical("Could not load a ServiceConfig, exiting...")
		return
	}

	log.Info("Host: " + cfg.Host)
}


