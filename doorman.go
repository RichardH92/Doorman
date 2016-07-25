package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/auth", HandleAuth).Queries("username", "{username:[0-9A-Za-z]+}", "password_hash", "{password_hash:[0-9A-Za-z]+}")

  log.Fatal(http.ListenAndServe(":8000", r))
}
