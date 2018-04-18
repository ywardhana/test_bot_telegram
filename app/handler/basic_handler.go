package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (*BasicHandler) HealthzHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	return nil
}

func (*BasicHandler) MetricHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	promhttp.Handler().ServeHTTP(w, r)
	return nil
}
