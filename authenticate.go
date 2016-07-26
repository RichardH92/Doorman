package main

import (
  "net/http"
  "fmt"
  "github.com/dvsekhvalnov/jose2go"
)

func HandleAuth(w http.ResponseWriter, req *http.Request) {
  if err := req.ParseForm(); err != nil {
     w.WriteHeader(401)
  }

  username := req.Form.Get("username")
  password := req.Form.Get("password")

  auth := authenticateUser(username, password)
  if auth == true {
    token, success := createToken(username)
    if success == false {
      w.WriteHeader(500)
    } else {
      w.Write([]byte("{\"token\":\"" + token + "\"}"))
    }
  } else {
    w.WriteHeader(401)
  }
}

func authenticateUser(username string, password string) bool {
  // TODO: Query here to validate the username and password

  return true
}

func createToken(username string) (string, bool) {
  payload := "{\"username\":\"" + username + "\"}"
  key := GetKey()

  token, err := jose.Sign(payload, jose.HS256, key)

  if(err == nil) {
      return token, true
  }

  fmt.Println("Error generating token.")
  return token, false
}
