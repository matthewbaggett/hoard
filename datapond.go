package main

import (
	_ "embed"
	"flag"
	log "github.com/sirupsen/logrus"
	"main/pkg/datapond"
)

func main() {
	var bindAddress string
	var bindPort int
	flag.StringVar(&bindAddress, "http-bind", "0.0.0.0", "Address to bind on")
	flag.IntVar(&bindPort, "http-port", 0, "Port to listen on")

	flag.Parse()

	handler, err := datapond.StartHandler(bindAddress, bindPort)

	log.Infof("Bind: %s:%d", bindAddress, bindPort)

	if err != nil {
		log.Errorf("Error starting Hoard DataPond: %v", err)
		return
	}

	log.Infof("Started Hoard DataPond listening on %s\n", handler.Listener.Addr())
}
