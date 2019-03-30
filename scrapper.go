package main

import (
//  "fmt"
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
        if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
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

func Scrape() Rooms {
  var rooms = GetLoc()
  var scrape = Rooms{}
  for _, room := range rooms {
    scrape = append(scrape, GetInfo(room))
  }

  return scrape
}

func String(n int32) string {
  buf := [11]byte{}
  pos := len(buf)
  i := int64(n)
  signed := i < 0
  if signed {
    i = -i
  }
  for {
    pos--
    buf[pos], i = '0' + byte(i % 10), i / 10
    if i == 0 {
      if signed {
        pos--
        buf[pos] = '-'
      }
      return string(buf[pos:])
    }
  }
}
