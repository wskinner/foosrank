package main

import (
	"github.com/kurrik/twittergo"
	"net/http"
	"os"
	"fmt"
)


func main() {
	var (
		err    error
		client *twittergo.Client
		req    *http.Request
		resp   *twittergo.APIResponse
		user   *twittergo.User
	)

	client, err = *twittergo.LoadCredentials()
	if err != nil {
		fmt.Printf("Could not parse CREDENTIALS file: %v\n", err)
		os.Exit(1)
	}
}
