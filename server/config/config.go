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

	// first we try to parse the config at the global configuration path
	if fileExists("/etc/festivals-fileserver.conf") {
		config := ParseConfig("/etc/festivals-fileserver.conf")
		if config != nil {
			return config
		}
	}

	// if there is no global configuration check the current folder for the template config file
	// this is mostly so the application will run in development environment
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("server initialize: could not read current path.")
	}
	path = path + "/config_template.toml"
	return ParseConfig(path)
}

func ParseConfig(cfgFile string) *Config {

	content, err := toml.LoadFile(cfgFile)
	if err != nil {
		log.Fatal("server initialize: could not read config file at '" + cfgFile + "'. Error: " + err.Error())
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

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
// see: https://golangcode.com/check-if-a-file-exists/
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
