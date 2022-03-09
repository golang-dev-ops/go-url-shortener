package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const EnvPrefix = "APP"

type Config interface {
	LoadConfig(configType string) (*config, error)
}

type config struct {
	App     App     `json:"app,omitempty"`
	Storage Storage `json:"storage,omitempty"`
	// Redis Redis `json:"redis,omitempty"`
	// Mysql MySQL `json:"mysql,omitempty"`
}

type App struct {
	Port int `json:"port,omitempty"`
}

type Storage struct {
	Redis Redis `json:"redis,omitempty"`
}

type Redis struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
}
type MySQL struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

func LoadConfig(configType string) (*config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType(configType)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshall config: %w", err)
	}
	return &cfg, nil
}
