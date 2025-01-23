package github

import (
	//"encoding/json"

	"io"
	"log"
	"net/http"
	"os"
)

func APIPractice() {
	resp, err := http.Get("https://api.github.com/users/dinesh-g1")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: can't copy - %v", err.Error())
	}
}
