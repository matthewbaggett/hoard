package hoard_api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"main/pkg/common"
	"net"
	"net/http"
)

type Handler struct {
	Router   *httprouter.Router
	Listener net.Listener
	Server   *http.Server
	Logger   logrus.FieldLogger
	Port     int
	Address  string
}

func (h Handler) PutEntityInHoard(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (h Handler) HealthCheck(writer http.ResponseWriter, request *http.Request) {
	h.Logger.Infof("Healthcheck Requested")

	h.responseJSON(writer, request, 200, map[string]any{
		"status":  "alive",
		"version": common.GetVersion(),
	})
	return
}
func (h *Handler) responseJSON(w http.ResponseWriter, r *http.Request, code int, v ...any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var data []byte
	if len(v) == 0 || v[0] == nil {
		data, _ = json.MarshalIndent(struct{}{}, "", "  ")
	} else if err, ok := v[0].(error); ok {
		h.Logger.Errorf("%v %v: %v", r.Method, r.RequestURI, err)
		data, _ = json.MarshalIndent(map[string]any{
			"error": err.Error(),
		}, "", "  ")
	} else {
		data, _ = json.MarshalIndent(v[0], "", "  ")
	}
	w.WriteHeader(code)
	_, _ = w.Write(data)
}

func (h Handler) GetAddress() string {
	return fmt.Sprintf("%s:%d", h.Address, h.Port)
}
