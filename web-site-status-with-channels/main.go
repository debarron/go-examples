package main

import (
  "fmt"
  "net/http"
  "time"
)

func main() {
  links := []string{
    "http://google.com",
    "http://facebook.com",
    "http://stackoverflow.com",
    "http://golang.org",
    "http://amazon.com",
  }

  myChannel := make(chan string)

  for _, link := range links{
    go checkLink(link, myChannel)
  }

  /*
  for i := 0; i < len(links); i++ {
    fmt.Println(<-myChannel)
  }
  
  for {
    go checkLink(<-myChannel, myChannel)
  }
  */

  for linkFromChan := range myChannel {
    go func(l string) {
      checkLink(l, myChannel)
      time.Sleep(5*time.Second)
    }(linkFromChan)
  }

}

func checkLink(link string, c chan string) {
  _, err := http.Get(link)
  if err != nil {
    c <- link
    fmt.Println("it appears might be down: ", link)
    return
  }

  fmt.Println(link, " is up!")
  c <- link
}
