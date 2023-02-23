package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Login    string
	Name     string
	NumRepos int `json:"public_repos"`
}

type GithubResponse struct {
	Name        string
	PublicRepos int
}

const githubUrlTemplate = "https://api.github.com/users/%s"

func userInfo(login string) (*User, error) {
	url := fmt.Sprintf(githubUrlTemplate, login)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: can't call %s", url)
	}
	defer resp.Body.Close()

	// Decode JSON
	user := User{Login: login}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error : %s", err)
	}

	fmt.Printf("%#v\n", user)
}
