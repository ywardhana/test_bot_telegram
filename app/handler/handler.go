package handler

import (
	"net/http"

	presponse "github.com/bukalapak/packen/response"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
}

type BasicHandler struct {
}

type AgentHandler struct {
}

func (*Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	ok_response(w, http.StatusOK)
	return nil
}

func ok_response(w http.ResponseWriter, httpCode int) {
	object_response(w, struct{ Message string }{"ok"}, httpCode)
}

func object_response(w http.ResponseWriter, o interface{}, httpCode int) {
	successResponse := presponse.BuildSuccess(o, presponse.MetaInfo{HTTPStatus: httpCode})
	presponse.Write(w, successResponse, 200)
}
