package main

import (
	_ "embed"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	hoard_api "main/pkg/api"
)

func main() {
	log.Println("Starting Hoard DataPond!")
	handler := &hoard_api.Handler{}
	httpRouter := httprouter.New()
	httpRouter.GET("/health", handler.HealthCheck)
	httpRouter.PUT("/hoard/:entity_type", handler.PutEntityInHoard)
}
