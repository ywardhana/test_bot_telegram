package handler

import (
  "encoding/json"
  "net/http"
  "strconv"
  "log"

  "github.com/julienschmidt/httprouter"
  "github.com/bukalapak/kingsman/app/utility"
  "github.com/bukalapak/kingsman/app/entity"
  "github.com/bukalapak/kingsman/app/service"
  "github.com/bukalapak/kingsman/db"
)

func (*AgentHandler) GetAgentHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
  err := utility.CheckBasicAuth(r)
  if err != nil {
    log.Println("error:" + err.Error())
    utility.SendErrorResponse(w, entity.AuthorizationFailedError)
    return nil
  }

  database, _ := db.InitMysql()

  id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
  if err != nil {
    log.Println("error:" + err.Error())
    utility.SendErrorResponse(w, entity.AgentNotFoundError)
    return nil
  }

  agent, customError := service.GetAgent(database, id)

  if customError.Error() != "" {
    log.Println("error:" + customError.Error())
    utility.SendErrorResponse(w, customError)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(&agent)
  return nil
}
