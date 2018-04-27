package utility

import (
  "encoding/json"
  "errors"
  "net/http"
  "os"
  "regexp"

  "github.com/bukalapak/aleppo/crypto/jose"
  "github.com/bukalapak/kingsman/app/entity"
)

// DecodeJoseToken decodes authorization token into BukalapakTokenInfo
func DecodeJoseToken(token string) (*entity.BukalapakTokenInfo, error) {
  match, _ := regexp.MatchString("^Token ", token)

  tokenInfo := &entity.BukalapakTokenInfo{}
  if !match {
    return tokenInfo, errors.New("Invalid Token")
  }
  // Strip the `Token ` part
  pToken := token[6:]

  raw, err := jose.New().Decode(pToken)
  if err != nil {
    return tokenInfo, err
  }

  err = json.Unmarshal(raw, tokenInfo)
  if err != nil {
    return tokenInfo, err
  }

  return tokenInfo, nil
}

func CheckBasicAuth(r *http.Request) error {
  headerUsername, headerPass, ok := r.BasicAuth()

  if !ok {
    return errors.New("Authorization failed")
  }

  username := os.Getenv("KINGSMAN_BASIC_USER")
  pass := os.Getenv("KINGSMAN_BASIC_PASSWORD")

  if username != headerUsername {
    return errors.New("Authorization failed")
  }

  if pass != headerPass {
    return errors.New("Authorization failed")
  }

  return nil
}
