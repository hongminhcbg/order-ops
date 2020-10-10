package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MySQLConnectURL string `envconfig:"MYSQL_CONNECT_URL" require:"true"`
}

func GetConfig() (*Config, error) {
	var result Config
	err := envconfig.Process("", &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
