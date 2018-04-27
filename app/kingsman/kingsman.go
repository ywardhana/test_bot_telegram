package kingsman

import (
	"net/http"

	"github.com/bukalapak/kingsman/app/handler"
	"github.com/bukalapak/packen/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/subosito/gotenv"
)

func NewKingsman() http.Handler {
	gotenv.Load()

	router := httprouter.New()

	h := handler.Handler{}
	basicHandler := handler.BasicHandler{}
	agentHandler := handler.AgentHandler{}
	router.GET("/", middleware.MonitorHTTP("index", h.Index))
	router.GET("/healthz", middleware.MonitorHTTP("healthz", basicHandler.HealthzHandler))
	router.GET("/metrics", middleware.MonitorHTTP("metrics", basicHandler.MetricHandler))
	router.GET("/agents/:id", middleware.MonitorHTTP("metrics", agentHandler.GetAgentHandler))
	return router
}
