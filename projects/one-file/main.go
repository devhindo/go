package main

import (
   "crypto/rand"
   "encoding/base64"
   "fmt"
)

func generateRandomString(length int) string {
   b := make([]byte, length)
   _, err := rand.Read(b)
   if err != nil {
      panic(err)
   }
   return base64.StdEncoding.EncodeToString(b)
}

func main() {
   randomString := generateRandomString(10)
   fmt.Println(randomString)
}