package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"thorns/config"
	"thorns/server"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := config.Init(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
