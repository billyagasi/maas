// config/config.go
package config

import "github.com/spf13/viper"

type Config struct {
	Elasticsearch struct {
		Host   string
		Port   int
		Scheme string
	}
}

func LoadConfig() (*Config, error) {
	var config Config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // or json, etc.

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	err := viper.Unmarshal(&config)
	return &config, err
}
