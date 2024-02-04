package main

import (
	_ "embed"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	hoardapi "main/pkg/api"
)

func main() {
	log.Println("Starting Hoard DataLake!")
	handler := &hoardapi.Handler{}
	httpRouter := httprouter.New()
	httpRouter.GET("/health", handler.HealthCheck)
}
