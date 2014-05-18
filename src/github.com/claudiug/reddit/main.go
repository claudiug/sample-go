package main

import (
  // "os"
  "net/http"
  "log"
  // "io"
  "fmt"
  "encoding/json"
)

type Item struct {
  Title string
  Url string
}

type Response struct {
  Data struct {
    Children []struct {
      Data Item
    }
  }
}

func main() {
  response, err := http.Get("http://reddit.com/r/golang.json")
  if err != nil {
    log.Fatal(err)
  }
  if response.StatusCode != http.StatusOK {
    log.Fatal(response.StatusCode)
  }
  // _, err = io.Copy(os.Stdout, response.Body)
  r := new(Response)
  err = json.NewDecoder(response.Body).Decode(r)
  if err != nil {
    log.Fatal(err)
  }

  for _, child := range r.Data.Children {
    fmt.Println(child.Data.Title)
  }
}

