package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"thorns/config"
	"thorns/server"
)

// @title Demo api
// @version 1.0
// @description Демонстрационное приложение

// @host 127.0.0.1:81
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := config.Init(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port"), viper.GetString("external_url")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
