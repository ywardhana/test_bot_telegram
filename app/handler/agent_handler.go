package handler

import (
  "encoding/json"
  "net/http"
  "log"

  "github.com/julienschmidt/httprouter"
  "github.com/bukalapak/kingsman/app/utility"
  "github.com/bukalapak/kingsman/app/entity"
)

func (*AgentHandler) GetAgentHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
  err := utility.CheckBasicAuth(r)
  if err != nil {
    log.Println("error:" + err.Error())
    utility.SendErrorResponse(w, entity.AuthorizationFailedError)
    return nil
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(&err)
  return nil
}
