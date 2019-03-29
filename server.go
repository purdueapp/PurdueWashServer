package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "log"
)

func main() {
  fmt.Println("Server started and running at port 8421")

  http.HandleFunc("/", req)
  log.Fatal(http.ListenAndServe(":8421", nil))
}

func req(w http.ResponseWriter, r *http.Request) {

  laundryRooms := Rooms{}
  laundryRooms = Scrape()

  js, err := json.Marshal(laundryRooms)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
