package config

import (
	"os"

	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ServiceBindAddress string
	ServiceBindHost    string
	ServicePort        int
	ServiceKey         string
	TLSRootCert        string
	TLSCert            string
	TLSKey             string
	LoversEar          string
	Interval           int
	APIKeys            []string
	AdminKeys          []string
	StorageURL         string
	ResizeStorageURL   string
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
		log.Fatal().Msg("server initialize: could not read default config file with error:" + err.Error())
	}
	path = path + "/config_template.toml"
	return ParseConfig(path)
}

func ParseConfig(cfgFile string) *Config {

	content, err := toml.LoadFile(cfgFile)
	if err != nil {
		log.Fatal().Msg("server initialize: could not read config file at '" + cfgFile + "' with error: " + err.Error())
	}

	serviceBindAdress := content.Get("service.bind-address").(string)
	serviceBindHost := content.Get("service.bind-host").(string)
	serverPort := content.Get("service.port").(int64)
	serviceKey := content.Get("service.key").(string)

	tlsrootcert := content.Get("tls.festivaslapp-root-ca").(string)
	tlscert := content.Get("tls.cert").(string)
	tlskey := content.Get("tls.key").(string)

	loversear := content.Get("heartbeat.endpoint").(string)
	interval := content.Get("heartbeat.interval").(int64)

	keyValues := content.Get("authentication.api-keys").([]interface{})
	keys := make([]string, len(keyValues))
	for i, v := range keyValues {
		keys[i] = v.(string)
	}
	adminKeyValues := content.Get("authentication.admin-keys").([]interface{})
	adminKeys := make([]string, len(adminKeyValues))
	for i, v := range adminKeyValues {
		adminKeys[i] = v.(string)
	}

	storageURL := content.Get("service.storage-url").(string)
	servicResizedStorageURL := content.Get("service.resized-storage-url").(string)

	return &Config{
		ServiceBindAddress: serviceBindAdress,
		ServiceBindHost:    serviceBindHost,
		ServicePort:        int(serverPort),
		ServiceKey:         serviceKey,
		TLSRootCert:        tlsrootcert,
		TLSCert:            tlscert,
		TLSKey:             tlskey,
		LoversEar:          loversear,
		Interval:           int(interval),
		APIKeys:            keys,
		AdminKeys:          adminKeys,
		StorageURL:         storageURL,
		ResizeStorageURL:   servicResizedStorageURL,
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
