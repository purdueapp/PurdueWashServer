package main

import (
  "net/http"
  "encoding/json"
  "fmt"
)

func req(w http.ResponseWriter, r *http.Request) {
  var lol, err = json.Marshal(getLoc())
  if err != nil {
    fmt.Fprint(w, lol)
  } else {
    fmt.Fprint(w, err)
  }
}

func main() {
  http.HandleFunc("/", req)
  http.ListenAndServe(":8000", nil)
}
