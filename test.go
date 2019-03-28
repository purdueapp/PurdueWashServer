package main

import (
  "fmt"
  "github.com/gocolly/colly"
)

var url = "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx"

func main() {
  c := colly.NewCollector()

  c.OnHTML("tr.MachineReadyMode", func(e *colly.HTMLElement) {
    fmt.Println(e.ChildText("td.name"))
  })

  c.Visit(url)
}
