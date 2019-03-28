package main

import (
  "fmt"
  "github.com/gocolly/colly"
)

func main() {
  c := colly.NewCollector()

  c.OnRequest(func(r *colly.Request) {
    fmt.Println("visiting", r.URL)
  })

  c.OnHTML("table tbody tr td center h2", func(e *colly.HTMLElement) {
    fmt.Println(e.Text)
  })

  c.OnHTML("table body tr td", func(e *colly.HTMLElement) {
    e.ForEach("header", func(_ int, el *colly.HTMLElement) {
        fmt.Println(el.Text)
    })
  })

  c.Visit("http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx")
}
