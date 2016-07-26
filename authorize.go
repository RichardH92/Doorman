package main

import (
  //"fmt"
  "encoding/json"
  "net/http"
  "github.com/dvsekhvalnov/jose2go"
)

func authorized(username string, endpoint string) bool {

  // TODO: Query here to see if the user is authorized to access endpoint

  return true
}

func HandleTest(w http.ResponseWriter, req *http.Request) {
  sharedKey := GetKey()

  if err := req.ParseForm(); err != nil {
    w.WriteHeader(401)
  }

  token := req.Form.Get("token")
  payload, _, err := jose.Decode(token, sharedKey)

  if(err==nil) {
    username := getUsernameFromJSON(payload)

    if authorized(username, "/test") {
      //TODO: Forward the request to the proper location
    }

  } else {
    w.WriteHeader(401)
  }
}

func getUsernameFromJSON(payload string) string {
  var data map[string]interface{}

  if err := json.Unmarshal([]byte(payload), &data); err != nil {
      panic(err)
  }

  return data["username"].(string)
}


// TODO: Write handlers for the rest of your API here
