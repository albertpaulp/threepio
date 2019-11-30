package main

import (
	"fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "os"
)

type Commit struct {
  Sha string `json:"sha"`
  Url string `json:"html_url"`
}

type GithubResponse []Commit

func exit(msg string){
  fmt.Println(msg)
  os.Exit(1)
}

func main() {
  client := &http.Client{}
  repo_name, exists := os.LookupEnv("GITHUB_REPO")
  if exists == false {
    exit("ENV:GITHUB_REPO not set !")
  }
  secret_token_env, exists := os.LookupEnv("GITHUB_TOKEN")
  githubToken := fmt.Sprintf("token %s", secret_token_env)
  if exists == false {
    exit("ENV:GITHUB_TOKEN not set !")
  }
  url := fmt.Sprintf("https://api.github.com/repos/%s/commits?since=2019-11-26T00:00:00", repo_name)
  request, err := http.NewRequest("GET", url, nil)
  request.Header.Add("Authorization", githubToken)
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

