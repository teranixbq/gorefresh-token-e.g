package config

import "refresh/pkg/validation"

type Config struct {
	DBHOST         string `env:"DBHOST"`
	DBPORT         int    `env:"DBPORT"`
	DBUSER         string `env:"DBUSER"`
	DBPASS         string `env:"DBPASS"`
	DBNAME         string `env:"DBNAME"`
	ACCES_TOKEN    string `env:"ACCES_TOKEN"`
	REFRESH_TOKEN  string `env:"REFRESH_TOKEN"`
	ACCES_EXPIRE   int    `env:"ACCES_EXPIRE"`
	REFRESH_EXPIRE int    `env:"REFRESH_EXPIRE"`
}

func InitConfig() *Config {
	config := &Config{}
	validation.EnvCheck(config)
	return config
}
