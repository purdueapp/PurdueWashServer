package main

import (
  "fmt"
  "github.com/gocolly/colly"
)

var url = "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx"

func main() {
  c := colly.NewCollector()

  c.OnRequest(func(r *colly.Request) {
    fmt.Println("visiting", r.URL)
  })

  //c.OnHTML("table tbody tr td center h2", func(e *colly.HTMLElement) {
  //  fmt.Println(e.Text)
  //})

  c.OnHTML("table body tr", func(e *colly.HTMLElement) {
    e.ForEach("", func(_ int, el *colly.HTMLElement) {
        fmt.Println(el.ChildText("td[header]"))
    })
  })

  c.Visit(url)
}
