package main

import (
  "net/http"
  "log"
  "github.com/gorilla/mux"
)

func main() {
  GenerateKey()
  r := mux.NewRouter()

  r.HandleFunc("/auth", HandleAuth).
  Queries("username", "{username:[0-9A-Za-z]+}", "password_hash", "{password_hash:[0-9A-Za-z]+}").
  Methods("POST")

  r.HandleFunc("/verify", HandleVerify).
  Queries("username", "{username:[0-9A-Za-z]+}").
  Methods("POST")

  log.Fatal(http.ListenAndServe(":8000", r))
}
