package utility

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/bukalapak/kingsman/app/entity"
  "github.com/bukalapak/kingsman/app/resource/response"
)

func SendSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
  response := response.SuccessResponse{
    Data: data,
    Meta: response.MetaInfo{
      HttpStatus: statusCode,
    },
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(response.Meta.HttpStatus)
  json.NewEncoder(w).Encode(&response)
}

func SendErrorResponse(w http.ResponseWriter, err entity.CustomError) {
  log.Println("error: " + err.Error())
  response := response.ErrorResponse{
    Errors: []response.ErrorInfo{
      response.ErrorInfo{
        Message: err.Error(),
        Code:    err.Code,
        Field:   err.Field,
      },
    },
    Meta: response.MetaInfo{
      HttpStatus: err.HttpStatus,
    },
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(response.Meta.HttpStatus)
  json.NewEncoder(w).Encode(&response)
}
