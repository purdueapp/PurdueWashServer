package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "github.com/bxcodec/faker"
)

type LaundryRoom struct {
  Name string
  Url string
  ImageUrl string
  AvailableWashers int32
  TotalWashers int32
  AvailableDryers int32
  TotalDryers int32
  Machines []Machine
}

type Machine struct {
  Name string
  Status string
  TimeRemaining int32
}

func req(w http.ResponseWriter, r *http.Request) {

  data := LaundryRoom{}
  faker.FakeData(&data)
  fmt.Println(json.Marshal(data))
  fmt.Fprint(w, data)
}

func main() {
  http.HandleFunc("/", req)
  http.ListenAndServe(":8421", nil)
}
