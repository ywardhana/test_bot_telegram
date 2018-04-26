package entity

import (
  "fmt"
  "net/http"
)

type CustomError struct {
  Message    string `json:"message"`
  Code       int    `json:"code"`
  Field      string `json:"field,omitempty"`
  HttpStatus int    `json:"http_status"`
}

func (e CustomError) Error() string {
  return fmt.Sprintf("%s", e.Message)
}

var (
  InvalidTokenError = CustomError{
    Message:    "Invalid token",
    Code:       88101,
    Field:      "",
    HttpStatus: http.StatusUnauthorized,
  }
  UserUnauthorizedError = CustomError{
    Message:    "User unauthorized",
    Code:       88102,
    Field:      "",
    HttpStatus: http.StatusForbidden,
  }
  FailedDecodeJSONError = CustomError{
    Message:    "Invalid parameters",
    Code:       88103,
    Field:      "",
    HttpStatus: http.StatusBadRequest,
  }
  AuthorizationFailedError = CustomError{
    Message:    "Authorization failed",
    Code:       88104,
    HttpStatus: http.StatusUnauthorized,
  }
)
