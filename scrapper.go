package main

import (
  "fmt"
  "strings"
  "github.com/gocolly/colly"
)

var url = "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx"

func GetLoc() Rooms {
  var rooms = Rooms{}

  c := colly.NewCollector()

  c.OnHTML("h2 a[href]", func(e *colly.HTMLElement) {
    room := Room{}
    room.Name = e.Text
    room.Url = e.Attr("href")
    rooms = append(rooms, room)
  })

  c.Visit(url)

  return rooms
}

func GetInfo(room Room) Room {
  var machines = Machines{}
  var availWash int32 = 0
  var availDry int32 = 0
  var wash int32 = 0
  var dry int32 = 0

  c := colly.NewCollector()
  c.OnHTML("tr.MachineReadyMode", func(e *colly.HTMLElement) {
    machine := Machine{}
    machine.Name = e.ChildText("td.name")
    machine.Status = e.ChildText("td.status")
    machine.TimeRemaining = e.ChildText("time")
    machines = append(machines, machine)

    if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
      availDry++
      dry++
    } else {
      availWash++
      wash++
    }
  })

  c.OnHTML("MachineRunMode", func(e *colly.HTMLElement) {
    machine := Machine{}
    machine.Name = e.ChildText("td.name")
    machine.Status = e.ChildText("td.status")
    machine.TimeRemaining = e.ChildText("time")
    machines = append(machines, machine)

    if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
      dry++
    } else {
      wash++
    }
  })

  c.OnHTML("MachineEndOfCycle", func(e *colly.HTMLElement) {
    machine := Machine{}
    machine.Name = e.ChildText("td.name")
    machine.Status = e.ChildText("td.status")
    machine.TimeRemaining = e.ChildText("time")
    machines = append(machines, machine)

    if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
      dry++
    } else {
      wash++
    }
  })

  c.OnHTML("MachineRunModeAlmostDone", func(e *colly.HTMLElement) {
    machine := Machine{}
    machine.Name = e.ChildText("td.name")
    machine.Status = e.ChildText("td.status")
    machine.TimeRemaining = e.ChildText("time")
    machines = append(machines, machine)

    if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
      dry++
    } else {
      wash++
    }
  })

  c.Visit(url + room.Url)

  room.AvailableWashers = availWash
  room.TotalWashers = wash
  room.AvailableDryers = availDry
  room.TotalDryers = dry
  room.Machines = machines

  return room
}

func Scrape() Rooms {
  var rooms = GetLoc()
  var scrape = Rooms{}
  for _, room := range rooms {
    scrape = append(scrape, GetInfo(room))
  }

  return scrape
}
