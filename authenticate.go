package main

import (
  "net/http"
  "fmt"
)

func HandleAuth(w http.ResponseWriter, req *http.Request) {
  if err := req.ParseForm(); err != nil {
     // handle error
     fmt.Println("Error")
  }
  username := req.Form.Get("username")
  password := req.Form.Get("passwod")

  fmt.Println(username)
  fmt.Println(password)

  w.Write([]byte("Auth\n" + username))
  w.Write([]byte("\n" + password))
}
