package main

import "net/http"
import "fmt"
import "io/ioutil"



func main() {

  url := "http://wpvitassuds01.itap.purdue.edu/washalertweb/washalertweb.aspx"
  resp, _ := http.Get(url)

  bytes, _ := ioutil.ReadAll(resp.Body)

  fmt.Println("HTML:\n\n", string(bytes))



  resp.Body.Close()

}
