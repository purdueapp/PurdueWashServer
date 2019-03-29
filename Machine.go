package main

type Machine struct {
  Name          string `json: "name"`
  Status        string `json: "status"`
  TimeRemaining string `json: "timeRemaining"`
}

type Machines []Machine
