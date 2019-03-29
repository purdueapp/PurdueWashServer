package main

import (
  "fmt"
  "github.com/gocolly/colly"
  "strings"
)

type Room struct {
  name string
  url string
  availableWashers int32
  totalWashers int32
  availableDryers int32
  totalDryers int32
  machines []Machine
}

type Machine struct {
  name string
  status string
  timeRemaining string
}

var url = "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx"

func GetLoc() []Room {
  var rooms = []Room{}

  c := colly.NewCollector()

  c.OnHTML("h2 a[href]", func(e *colly.HTMLElement) {
    temp := Room{}
    temp.name = e.Text
    temp.url = e.Attr("href")
    rooms = append(rooms, temp)
  })

  c.Visit(url)

  return rooms
}

func GetInfo(dorm Room) Room {
  var machines = []Machine{}
  var availWash int32 = 0
  var availDry int32 = 0
  var wash int32 = 0
  var dry int32 = 0

  c := colly.NewCollector()
  c.OnHTML("tr.MachineReadyMode", func(e *colly.HTMLElement) {
    temp := Machine{}
    temp.name = e.ChildText("td.name")
    temp.status = e.ChildText("td.status")
    temp.timeRemaining = e.ChildText("time")
    machines = append(machines, temp)

    if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
      availDry++
      dry++
    } else {
      availWash++
      wash++
    }
  })

  c.OnHTML("MachineRunMode", func(e *colly.HTMLElement) {
    temp := Machine{}
    temp.name = e.ChildText("td.name")
    temp.status = e.ChildText("td.status")
    temp.timeRemaining = e.ChildText("time")
    machines = append(machines, temp)

    if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
      dry++
    } else {
      wash++
    }
  })

  c.OnHTML("MachineEndOfCycle", func(e *colly.HTMLElement) {
    temp := Machine{}
    temp.name = e.ChildText("td.name")
    temp.status = e.ChildText("td.status")
    temp.timeRemaining = e.ChildText("time")
    machines = append(machines, temp)

    if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
      dry++
    } else {
      wash++
    }
  })

  c.OnHTML("MachineRunModeAlmostDone", func(e *colly.HTMLElement) {
    temp := Machine{}
    temp.name = e.ChildText("td.name")
    temp.status = e.ChildText("td.status")
    temp.timeRemaining = e.ChildText("time")
    machines = append(machines, temp)

    if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
      dry++
    } else {
      wash++
    }
  })

  c.Visit(url + dorm.url)

  dorm.availableWashers = availWash
  dorm.totalWashers = wash
  dorm.availableDryers = availDry
  dorm.totalDryers = dry
  dorm.machines = machines

  return dorm
}

func scrape() []Room {
  var rooms = GetLoc()
  var scrape = []Room{}
  for _, room := range rooms {
    scrape = append(scrape, GetInfo(room))
  }
  fmt.Println(scrape)
  return scrape
}
