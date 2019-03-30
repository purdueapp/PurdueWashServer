package main

type Room struct {
  Name              string    `json: "name"`
  Url               string    `json: "url"`
  AvailableWashers  int32     `json: "availableWashers"`
  TotalWashers      int32     `json: "totalWashers"`
  AvailableDryers   int32     `json: "availableDryers"`
  TotalDryers       int32     `json: "totalDryers"`
  Machines          []Machine `json: "machines"`
}

type Rooms []Room
