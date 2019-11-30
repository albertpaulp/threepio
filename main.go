package main

import (
	"fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Commit struct {
  Sha string `json:"sha"`
  Url string `json:"html_url"`
}

type GithubResponse []Commit

func main() {
  client := &http.Client{}
  url := ""
  request, err := http.NewRequest("GET", url, nil)
  request.Header.Add("Authorization", "token ")
  request.Header.Add("Accept", "application/vnd.github.v3+json")
  response, err := client.Do(request)
  if err != nil {
    fmt.Println("Errored out !!")
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  vals := GithubResponse{}
  err = json.Unmarshal(body, &vals)
  if err != nil {
    fmt.Println("Parsing errored out !")
    fmt.Println(err)
  }
  fmt.Println(vals)
}

