package config

import (
	"github.com/pelletier/go-toml"
	"log"
	"os"
)

type Config struct {
	StorageURL       string
	ResizeStorageURL string
	ServicePort      int
}

func DefaultConfig() *Config {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal("server initialize: could not read default config file")
	}
	path = path + "/config_template.toml"
	return ParseConfig(path)
}

func ParseConfig(cfgFile string) *Config {

	content, err := toml.LoadFile(cfgFile)
	if err != nil {
		log.Fatal("server initialize: could not read config file")
	}

	storage_url := content.Get("service.storage-url").(string)
	servic_resized_storage_url := content.Get("service.resized-storage-url").(string)
	serverPort := content.Get("service.port").(int64)

	return &Config{
		StorageURL:       storage_url,
		ResizeStorageURL: servic_resized_storage_url,
		ServicePort:      int(serverPort),
	}
}
