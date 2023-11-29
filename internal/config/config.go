package config

import "github.com/spf13/viper"

type Config struct {
	Elasticsearch struct {
		Host   string
		Port   int
		Scheme string
	}
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
