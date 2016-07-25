package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/gorilla/mux"
  "github.com/dvsekhvalnov/jose2go"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/auth", HandleAuth).
  Queries("username", "{username:[0-9A-Za-z]+}", "password_hash", "{password_hash:[0-9A-Za-z]+}").
  Methods("POST")

  r.HandleFunc("/test", HandleTest)

  log.Fatal(http.ListenAndServe(":8000", r))
}

func HandleTest(w http.ResponseWriter, req *http.Request) {
  if err := req.ParseForm(); err != nil {
    w.WriteHeader(401)
  }

  token := req.Form.Get("token")
  sharedKey :=[]byte{97,48,97,50,97,98,100,56,45,54,49,54,50,45,52,49,99,51,45,56,51,100,54,45,49,99,102,53,53,57,98,52,54,97,102,99}
  payload, headers, err := jose.Decode(token,sharedKey)

  if(err==nil) {
    //go use token
    fmt.Printf("\npayload = %v\n",payload)

    //and/or use headers
    fmt.Printf("\nheaders = %v\n",headers)
  } else {
    w.WriteHeader(401)
  }
}
