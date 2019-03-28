package main

import (
  "fmt"
  "github.com/gocolly/colly"
)

type room struct {
  name string
  imageUrl string
  availableWashers int32
  totalWashers int32
  availableDryers int32
  totalDryers int32
  machines []machine
}

type machine struct {
  name string
  status string
  timeRemaining int32
}

func main() {
  rooms := []room{}

  c := colly.NewCollector()

  c.OnHTML("table tbody tr td center", func(e *colly.HTMLElement) {
    e.ForEach("h2", func(_ int, el *colly.HTMLElement) {
      room := room{}
      room.name = e.Text
      rooms.append(rooms, room)
      fmt.Println(e.Text)

    })
  })

  c.Visit("http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx")
}
