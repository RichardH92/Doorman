package main

import (
  "crypto/rand"
  "fmt"
)

var key []byte

func GenerateKey() {
  key = make([]byte, 26)

  _, err := rand.Read(key)
  if err != nil {
    fmt.Println("Error generating key.")
  }
}

func GetKey() []byte {
  return key
}
