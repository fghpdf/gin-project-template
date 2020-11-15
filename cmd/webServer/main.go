package main

import (
	"fghpdf.me/gin-project-template/cmd/webServer/router"
	"fghpdf.me/gin-project-template/internal/pkg/config"
	"fghpdf.me/gin-project-template/internal/pkg/logger"
	"fghpdf.me/gin-project-template/internal/server/connectivity"
	log "github.com/sirupsen/logrus"
)

func main() {
	// init modules
	config.Init()
	logger.Init()
	app := router.Init()

	routers := app.Group("/api")

	routers.GET("/ping", connectivity.Ping)

	err := app.Run()
	if err != nil {
		log.Fatalf("App start err: %v", err)
	}
}
