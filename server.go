package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "log"
)

type LaundryRoom struct {
        Name                string    `json:"name"`
        Url                 string    `json:"url"`
        AvailableWashers    int32     `json:"availableWashers"`
        TotalWashers        int32     `json:"totalWashers"`
        AvailableDryers     int32     `json:"availableDryers"`
        TotalDryers         int32     `json:"totalDryers"`
        Machines            []Machine `json:"machines"`
}

type Machine struct {
        Name          string  `json:"name"`
        Status        string  `json:"status"`
        TimeRemaining int32   `json:"timeRemaining"`
}

func main() {
  fmt.Println("Server started and running at port 8421")

  http.HandleFunc("/", req)
  log.Fatal(http.ListenAndServe(":8421", nil))
}

func req(w http.ResponseWriter, r *http.Request) {

  data := LaundryRoom{}
  data.Name = "Wiley"
  data.AvailableWashers = 43
  data.TotalWashers = 23
  data.AvailableDryers = 53
  data.TotalDryers = 12

  machines := []Machine{}
  machines = append(machines, Machine{"Dryer 01", "out of order", 0})
  machines = append(machines, Machine{"Dryer 02", "available", 0})
  machines = append(machines, Machine{"Dryer 03", "in use", 5})
  machines = append(machines, Machine{"Dryer 04", "end of cycle", 0})
  machines = append(machines, Machine{"Dryer 05", "out of order", 0})
  machines = append(machines, Machine{"Dryer 06", "out of order", 0})

  data.Machines = machines
  laundryRooms := [1]LaundryRoom{data}

  js, err := json.Marshal(laundryRooms)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
