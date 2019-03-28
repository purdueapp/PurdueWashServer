package main

import (
  "fmt"
  "github.com/gocolly/colly"
)

type room struct {
  Name string
  Url string
  ImageUrl string
  AvailableWashers int32
  TotalWashers int32
  AvailableDryers int32
  TotalDryers int32
  Machines []machine
}

type machine struct {
  Name string
  Status string
  TimeRemaining int32
}

var url = "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx"

func getLoc() []room {
  var rooms = []room{}

  c := colly.NewCollector()

  c.OnHTML("h2 a[href]", func(e *colly.HTMLElement) {
    temp := room{}
    temp.Name = e.Text
    temp.Url = e.Attr("href")
    rooms = append(rooms, temp)
  })

  c.Visit(url)

  return rooms
}

func getMachines(dorm room) []machine {
  var machines = []machine{}

  c := colly.NewCollector()
  c.OnHTML("tr.MachineReadyMode", func(e *colly.HTMLElement) {
    temp := machine{}
    temp.Name = e.ChildText("td.name")
    temp.Status = e.ChildText("td.status")
    //temp.TimeRemaining = e.ChildText("time")
    machines = append(machines, temp)
  })

  c.OnHTML("MachineRunMode", func(e *colly.HTMLElement) {
  })

  c.OnHTML("MachineEndOfCycle", func(e *colly.HTMLElement) {
  })

  c.OnHTML("MachineRunModeAlmostDone", func(e *colly.HTMLElement) {
  })

  c.Visit(url + dorm.Url)

  fmt.Println(machines)

  return machines
}

func main() {
  var rooms = getLoc()
  for _, room := range rooms {
    room.Machines = getMachines(room)
  }

  fmt.Println(rooms)
}
