package main

import (
  "net/http"
  "fmt"
)

func req(w http.ResponseWriter, r *http.Request) {

  var lol string = getHTML()
  fmt.Fprintf(w, lol)

}

func main() {

  http.HandleFunc("/", req)
  http.ListenAndServe(":8000", nil)

}
