package kingsman

import (
	"net/http"

	"github.com/bukalapak/kingsman/app/handler"
	"github.com/bukalapak/packen/middleware"
	"github.com/julienschmidt/httprouter"
)

func NewKingsman() http.Handler {
	router := httprouter.New()

	h := handler.Handler{}
	basicHandler := handler.BasicHandler{}
	router.GET("/", middleware.MonitorHTTP("index", h.Index))
	router.GET("/healthz", middleware.MonitorHTTP("healthz", basicHandler.HealthzHandler))
	router.GET("/metrics", middleware.MonitorHTTP("metrics", basicHandler.MetricHandler))
	return router
}
