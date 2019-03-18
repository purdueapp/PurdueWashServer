package main

import (
  "net/http"
  "fmt"
)

func req(w http.ResponseWriter, r *http.Request) {

  fmt.Fprintf(w, "hi")

}

func main() {

  http.HandleFunc("/", req)
  http.ListenAndServe(":8000", nil)

}
