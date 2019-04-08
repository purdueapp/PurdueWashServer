package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "log"
  "github.com/gorilla/mux"
)

func main() {
  fmt.Println("Server started and running at port 8421")

  r := mux.NewRouter()
  r.HandleFunc("/{room}", RoomHandler).Queries("location", "{link}")
  r.HandleFunc("/", req)
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe(":8421", nil))
}

func req(w http.ResponseWriter, r *http.Request) {

  laundryRooms := Rooms{}
  laundryRooms = Scrape()

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(laundryRooms)
}

func RoomHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  room := Room {
    Name: vars["room"],
    Url: "?location=" + vars["link"],
  }
  room = GetInfo(room)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(room)
}
