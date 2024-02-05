package main

import (
	_ "embed"
	log "github.com/sirupsen/logrus"
	"main/pkg/datapond"
)

func main() {
	handler, err := datapond.StartHandler()

	if err != nil {
		log.Errorf("Error starting Hoard DataPond: %v", err)
		return
	}

	log.Infof("Started Hoard DataPond listening on %s\n", handler.Listener.Addr())
}
