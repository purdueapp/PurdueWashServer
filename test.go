package main

import (
  "fmt"
  "github.com/gocolly/colly"
  "strings"
)

var url = "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx"

func main() {
  c := colly.NewCollector()

//  c.OnHTML("tr.MachineReadyMode", func(e *colly.HTMLElement) {
//    fmt.Println(e.ChildText("td.name"))
//    fmt.Println(e.ChildText("td.status"))
//    fmt.Println(e.ChildText("td.time"))
//  })

  c.OnHTML("tr.MachineRunMode", func(e *colly.HTMLElement) {
    //fmt.Println(e.ChildText("td.name"))
    //fmt.Println(e.ChildText("td.status"))
    //fmt.Println(e.ChildText("td.time"))

    if strings.Compare(e.ChildText("td.type"), "Dryer") == 0 {
      fmt.Println("lol")
    }
  })


  c.Visit(url)
}
