package main

import (
  "fmt"
  "github.com/gocolly/colly"
  "strings"
)

//var url = "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx"
//var wiley = "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx?location=c29eba8b-63d1-4090-bd32-ea85c67f483c"

func main() {
  rooms := GetLoc()
  scrape := Rooms{}
  for _, room := range rooms {
    scrape = append(scrape, asdf(room))
  }

  fmt.Println(scrape)
}

func asdf(room Room) Room {
  c := colly.NewCollector()

  var machines = Machines{}
  var availWash int32 = 0
  var availDry int32 = 0
  var wash int32 = 0
  var dry int32 = 0

  c.OnHTML("tr", func(e *colly.HTMLElement) {
    test := e.ChildText("td.name")
    if strings.Compare(test, "") != 0 {
      machine := Machine{}
      machine.Name = e.ChildText("td.name")
      machine.Status = e.ChildText("td.status")
      machine.TimeRemaining = e.ChildText("time")
      machines = append(machines, machine)

      if strings.Compare(machine.Status, "Available") == 0 {
        if strings.Compare(e.ChildText("type"), "Dryer") == 0 {
          availDry++
        } else {
          availWash++
        }
      }

      dry++
      wash++
    }
  })

  c.Visit(url + room.Url)

  room.AvailableWashers = String(availWash)
  room.TotalWashers     = String(wash)
  room.AvailableDryers  = String(availDry)
  room.TotalDryers      = String(dry)
  room.Machines = machines

  return room
}
