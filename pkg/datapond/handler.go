package datapond

import (
	_ "embed"
	_ "fmt"
	log "github.com/sirupsen/logrus"
	hoardApi "main/pkg/api"
	"main/pkg/common"
	"net/http"
)

func StartHandler(bindAddress string, bindPort int) (*hoardApi.Handler, error) {

	log.Printf(
		"Starting Hoard DataPond! Version %s, built %s: %s\n",
		common.GetVersion(),
		common.GetBuildTime(),
		common.GetGitCommitMessage(),
	)
	handler := &hoardApi.Handler{
		Port:    bindPort,
		Address: bindAddress,
		Logger:  log.StandardLogger(),
	}

	http.HandleFunc("/health", handler.HealthCheck)
	http.ListenAndServe(handler.GetAddress(), nil)

	return handler, nil
}
