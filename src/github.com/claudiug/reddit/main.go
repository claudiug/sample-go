package main

import (
  // "os"
  "net/http"
  "log"
  // "io"
  "fmt"
  "encoding/json"
  "errors"
)

type Item struct {
  Title string
  Url string
  Comments int `json:"num_comments"`
  // num_comments int reddit name
}

type Response struct {
  Data struct {
    Children []struct {
      Data Item
    }
  }
}

func (i Item) String() string {
  comment := ""
  switch i.Comments{
  case 0:
    comment = "no comments"
  case 1:
    comment = "(1 comment)"
  default:
    comment = fmt.Sprintf("(%d comments)", i.Comments)
  }
  return fmt.Sprintf("%s\n\n%s\t%s", i.Title, comment, i.Url)
}

func get(reddit string)([] Item, error) {
  url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
  response, err := http.Get(url)
  if err != nil {
    // log.Fatal(err)
    return nil, err
  }
  defer response.Body.Close()
  if response.StatusCode != http.StatusOK {
    return nil, errors.New(response.Status)
  }
  r := new(Response)
  err = json.NewDecoder(response.Body).Decode(r)
  if err != nil {
    return nil, err
  }
  items := make([]Item, len(r.Data.Children))
  for i, child := range r.Data.Children {
    items[i] = child.Data
  }
  return items, nil
}

func main() {
  // response, err := http.Get("http://reddit.com/r/golang.json")
  // if err != nil {
  //   log.Fatal(err)
  // }
  // if response.StatusCode != http.StatusOK {
  //   log.Fatal(response.StatusCode)
  // }
  // // _, err = io.Copy(os.Stdout, response.Body)
  // r := new(Response)
  // decoded = json.NewDecoder(response.Body).Decode(r)
  // if decoded != nil {
  //   log.Fatal(decoded)
  // }

  // for _, child := range r.Data.Children {
  //   fmt.Println(child.Data.Title)
  // }


  items, err := get("ruby")
  if err != nil {
    log.Fatal(err)
  }
  for _, item := range items {
    fmt.Println(item)
  }
}

