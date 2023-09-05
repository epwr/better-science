// Defines a ServiceConfig struct, and parses TOML files from the `/config` directory to 
// populate the struct. Will load a `config.toml` if it exists, or a `local.toml` if it
// doesn't.
//
// Copywrite 2023 Eric Power - All Rights Reserved
package config

import (
	"github.com/BurntSushi/toml"
	"github.com/epwr/better-science/internal/common/eplog"
	"path/filepath"
	"runtime"
)

// ServiceConfig provides a central definition of all service-level configuration
// values & their types.
type ServiceConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`	
	KGEngine struct {
		PathSearch string `toml:"path_search"`
	} `toml:"kg_engine"`
}


// LoadConfig loads a config TOML file from the `/config` directory. If available,
// it will load the `config.toml` file. If that file is not available, or can not
// be appropriately parsed, then it will try to load `local.toml`.
//
// This function returns an error when neither `config.toml` or `local.toml` can
// be found & parsed.
func LoadConfig() (*ServiceConfig, error) {

	// Get the current file's path and directory
	_, currentFile, _, _ := runtime.Caller(0)
	
	// Navigate back to repository root
	newPath := currentFile
	for i := 0; i < 4; i++ {
		newPath = filepath.Dir(newPath)
	}
	configDir := filepath.Join(newPath, "config")

	var config ServiceConfig
	
	config_filename := filepath.Join(configDir, "config.toml")
	if err := load_config_file(&config, config_filename); err != nil {

		config_filename = filepath.Join(configDir, "local.toml")
		if err := load_config_file(&config, config_filename); err != nil {
			return nil, err
		}
	}
	
	return &config, nil
}

// load_config_file loads a configuration file into the provided ServiceConfig.
func load_config_file(service_config *ServiceConfig, filename string) (error) {

	if _, err := toml.DecodeFile(filename, &service_config); err != nil {
		log.Error(
			"Error in loading file '" + filename + "'.",
			"error_message", err.Error(),
		)
		return err
	}
	return nil		
}

