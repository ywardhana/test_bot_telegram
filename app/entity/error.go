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
    Code:       67101,
    Field:      "",
    HttpStatus: http.StatusUnauthorized,
  }
  UserUnauthorizedError = CustomError{
    Message:    "User unauthorized",
    Code:       67102,
    Field:      "",
    HttpStatus: http.StatusForbidden,
  }
  FailedDecodeJSONError = CustomError{
    Message:    "Invalid parameters",
    Code:       67103,
    Field:      "",
    HttpStatus: http.StatusBadRequest,
  }
  AuthorizationFailedError = CustomError{
    Message:    "Authorization failed",
    Code:       67104,
    HttpStatus: http.StatusUnauthorized,
  }

  TransactionNotFoundError = CustomError{
    Message:    "Transaction not found",
    Code:       67111,
    HttpStatus: http.StatusNotFound,
  }
  FeedbackNotFoundError = CustomError{
    Message:    "Feedback not found",
    Code:       67112,
    HttpStatus: http.StatusNotFound,
  }
  SellerNotFoundError = CustomError{
    Message:    "Seller not found",
    Code:       67113,
    HttpStatus: http.StatusBadRequest,
  }
  SellerNotConfirmedError = CustomError{
    Message:    "Seller not confirmed",
    Code:       67120,
    HttpStatus: http.StatusBadRequest,
  }
  PayOwnAccountError = CustomError{
    Message:    "Cannot pay your own account",
    Code:       67121,
    HttpStatus: http.StatusBadRequest,
  }
  InvalidAmountError = CustomError{
    Message:    "Transaction amount should be between 100 and 1000000",
    Code:       67122,
    HttpStatus: http.StatusBadRequest,
  }

  MothershipTransactionError = CustomError{
    Message:    "Failed to create hybrid transaction",
    Code:       67116,
    HttpStatus: http.StatusBadRequest,
  }
  TransactionUpdateStateError = CustomError{
    Message:    "Failed to change state",
    Code:       67115,
    HttpStatus: http.StatusBadRequest,
  }
  UnsupportedStateError = CustomError{
    Message:    "Unsupported state",
    Code:       67114,
    HttpStatus: http.StatusBadRequest,
  }
  StateTransitionError = CustomError{
    Message:    "Unable to change state",
    Code:       67118,
    HttpStatus: http.StatusBadRequest,
  }

  FeedbackAlreadyExistError = CustomError{
    Message:    "Feedback already exist",
    Code:       67119,
    HttpStatus: http.StatusNotAcceptable,
  }
  FeedbackNotEditableError = CustomError{
    Message:    "Feedback not editable",
    Code:       67117,
    HttpStatus: http.StatusBadRequest,
  }
  ModifyFeedbackError = CustomError{
    Message:    "Failed to change feedback",
    Code:       67203,
    HttpStatus: http.StatusBadRequest,
  }
  TransactionConfirmError = CustomError{
    Message:    "Failed to confirm transaction",
    Code:       67204,
    HttpStatus: http.StatusBadRequest,
  }
)
