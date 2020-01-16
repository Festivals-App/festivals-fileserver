package config

type Config struct {
	StorageURL       string
	ResizeStorageURL string
}

func GetConfig() *Config {
	return &Config{
		StorageURL:       "/srv/fileserver/images",
		ResizeStorageURL: "/srv/fileserver/images/resized",
	}
}
