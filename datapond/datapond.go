package main

import (
	_ "embed"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"main/pkg/datapond"
)

type datapond_connection struct {
	DSN      string
	Host     string
	Port     int
	Username string
	Password string
	Database string
}
type datapond_pool struct {
	IsConfigured bool
	Name         string
	Connections  []datapond_connection
	Filters      []string
}

// to string method for datapond_pool
func (dp datapond_pool) String() string {
	return fmt.Sprintf("DataPond Pool Named %s has %d Connections and %d Filters", dp.Name, len(dp.Connections), len(dp.Filters))
}

func main() {
	var bindAddress string
	var bindPort int
	var datapondPools []datapond_pool

	flag.StringVar(&bindAddress, "http-bind", "0.0.0.0", "Address to bind on")
	flag.IntVar(&bindPort, "http-port", 0, "Port to listen on")
	log.Infof("Creating flags for datapond pools")
	for i := 0; i < 10; i++ {
		datapondPools = append(datapondPools, datapond_pool{})
		//log.Infof("Creating flags for datapond pools: --pool%d-name", i)
		flag.StringVar(&datapondPools[i].Name, fmt.Sprintf("pool%d-name", i), "pool", "Name of the pool")
		for j := 0; j < 10; j++ {
			datapondPools[i].Connections = append(datapondPools[i].Connections, datapond_connection{})
			//log.Infof("Creating flags for datapond pools: --pool%d-dsn%d", i, j)
			flag.StringVar(&datapondPools[i].Connections[j].DSN, fmt.Sprintf("pool%d-dsn%d", i, j), "", "DSN for the connection")
		}
		datapondPools[i].Filters = append(datapondPools[i].Filters, "")
		flag.StringVar(&datapondPools[i].Filters[0], fmt.Sprintf("pool%d-filter", i), "", "Filter for the pool")
	}
	flag.Parse()

	// foreach datapond pool, if the name is set and atleast one dsn is set, mark it as configured
	for i := 0; i < len(datapondPools); i++ {
		if datapondPools[i].Name != "" {
			log.Infof("Checking pool: %s", datapondPools[i].String())
			for j := 0; j < len(datapondPools[i].Connections); j++ {
				if datapondPools[i].Connections[j].DSN != "" {
					datapondPools[i].IsConfigured = true
					log.Infof("Configured pool: %s", datapondPools[i].String())
					break
				}
			}
		}
	}

	handler, err := datapond.StartHandler(bindAddress, bindPort)

	log.Infof("Bind: %s:%d", bindAddress, bindPort)

	if err != nil {
		log.Errorf("Error starting Hoard DataPond: %v", err)
		return
	}

	log.Infof("Started Hoard DataPond listening on %s\n", handler.Listener.Addr())
}
