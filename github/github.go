package github

import (
	//"encoding/json"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Reply struct {
	Name        string
	PublicRepos int `json:"public_repos"`
}

func APIPractice() {
	url := "https://api.github.com/users/dinesh-g1"
	name, repos, err := githubInfo(url)
	if err != nil {
		log.Fatalf("error %v occurred", err)
	}
	fmt.Println(name, repos)
}

func githubInfo(login string) (string, int, error) {
	resp, err := http.Get(login)
	if err != nil {
		return "", 0, err
	}
	var rep Reply
	if resp.StatusCode == http.StatusOK {
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&rep)
		if err != nil {
			return "", 0, err
		}
		return rep.Name, rep.PublicRepos, nil
	} else {
		return "", 0, err
	}
}
