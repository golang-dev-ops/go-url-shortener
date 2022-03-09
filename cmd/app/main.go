package main

import (
	_ "encoding/json"
	_ "io/ioutil"
	_ "log"
	_ "net/http"
	_ "time"

	"github.com/golang-dev-ops/go-url-shortener/config"

	"go.uber.org/zap"
)

const (
	configType = "toml"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := config.LoadConfig(configType)
	if err != nil {
		logger.Fatal("Cannot load configuration from "+configType+", shutting down server...", zap.Error(err))
	}
	logger.Info("Config loaded", zap.Any("config", cfg))

	// TODO: Coding
}
