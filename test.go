package main

import (
  "fmt"
  "github.com/gocolly/colly"
)

var url = "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx"

func main() {
  c := colly.NewCollector()

//  c.OnHTML("tr.MachineReadyMode", func(e *colly.HTMLElement) {
//    fmt.Println(e.ChildText("td.name"))
//    fmt.Println(e.ChildText("td.status"))
//    fmt.Println(e.ChildText("td.time"))
//  })

  var count int32 = 0
  c.OnHTML("tr.MachineRunMode", func(e *colly.HTMLElement) {
    fmt.Println(e.ChildText("td.name"))
    fmt.Println(e.ChildText("td.status"))
    fmt.Println(e.ChildText("td.time"))
    count++
  })

  c.Visit(url)
  fmt.Println(count)
}
