package datapond

import (
	_ "embed"
	"errors"
	_ "fmt"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	hoard_api "main/pkg/api"
	"main/pkg/common"
	"net"
	"net/http"
	"time"
)

type Handler struct {
	Router   *httprouter.Router
	listener net.Listener
	server   *http.Server
	logger   log.FieldLogger
}

func StartHandler() (*hoard_api.Handler, error) {

	log.Printf(
		"Starting Hoard DataPond! Version %s, built %s\n",
		common.GetVersion(),
		common.GetBuildTime(),
	)
	handler := &hoard_api.Handler{}
	httpRouter := httprouter.New()
	httpRouter.GET("/health", handler.HealthCheck)
	httpRouter.PUT("/hoard/:entity_type", handler.PutEntityInHoard)
	handler.Router = httpRouter

	listener, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		return nil, err
	}
	server := &http.Server{
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           handler.Router,
	}
	go func() {
		if err := server.Serve(listener); err != nil && errors.Is(err, net.ErrClosed) {
			log.Errorf("http serve: %v", err)
		}
	}()
	handler.Listener = listener
	handler.Server = server

	log.Printf("Started Hoard DataPond on %s!\n", handler.Listener.Addr())

	return handler, nil
}
