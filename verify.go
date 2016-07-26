package main

import (
  "encoding/json"
  "net/http"
  "github.com/dvsekhvalnov/jose2go"
  "fmt"
)

func HandleVerify(w http.ResponseWriter, req *http.Request) {
  sharedKey := GetKey()

  if err := req.ParseForm(); err != nil {
     w.WriteHeader(500)
  }

  usernameFromQuery := req.Form.Get("username")
  token := req.Form.Get("token")
  payload, _, err := jose.Decode(token, sharedKey)

  if (err == nil) {
    usernameFromToken, unameErr := getUsernameFromJSON(payload)

    if (unameErr == true) {

      if (usernameFromToken == usernameFromQuery) {
        w.WriteHeader(200)
      } else {
        w.WriteHeader(401)
      }
    } else {
      w.WriteHeader(401)
    }

  } else {
    w.WriteHeader(401)
  }
}

func getUsernameFromJSON(payload string) (string, bool) {
  var data map[string]interface{}

  if err := json.Unmarshal([]byte(payload), &data); err != nil {
      fmt.Println("Error parsing username from the token's payload.")
      return "", false
  }

  return data["username"].(string), true
}
