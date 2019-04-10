package main

type Room struct {
  Name              string    `json: "name"`
  Url               string    `json: "url"`
  AvailableWashers  string    `json: "availableWashers"`
  TotalWashers      string    `json: "totalWashers"`
  AvailableDryers   string    `json: "availableDryers"`
  TotalDryers       string    `json: "totalDryers"`
  Machines          []Machine `json: "machines"`
}

type Rooms []Room
