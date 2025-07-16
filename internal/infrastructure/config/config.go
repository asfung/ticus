package config

import "github.com/jinzhu/configor"

type AppConfig struct {
	Environment string
	Port        string
	Database    DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewAppConfig(path string) (*AppConfig, error) {
	config := new(AppConfig)
	err := configor.
		New(&configor.Config{ErrorOnUnmatchedKeys: true}).
		Load(config, path)
	if err != nil {
		return nil, err
	}
	return config, nil
}
