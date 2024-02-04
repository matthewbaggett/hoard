package Hoard

import (
	"Hoard/pkg/api"
	_ "embed"
	"github.com/julienschmidt/httprouter"
)

func main() {
	handler := &api.Handler{}
	httpRouter := httprouter.New()
	httpRouter.GET("/health", handler.HealthCheck)
}
