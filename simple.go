package main

import (
	"github.com/wskinner/anaconda"
	"fmt"
	"net/url"
)

func getApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("WyWJsRBoWWQvnG8jnoO0bA")
	anaconda.SetConsumerSecret("1QxBiTHqcy5jeMgFzxwJyon5tdThphRzzOlmxiW3qG8")
	var api *anaconda.TwitterApi = anaconda.NewTwitterApi("2295398977-RCw9S9GM5UZ2lkQC3prlNfAKWgnaUdoRdcBEh6k", "olpgzfjU4a47f67kJQh7TrZ5ZjvOVE1RLUaD0RpbyD4JH")

	return api
}

func main() {

	api := getApi()

	v := url.Values{}
	mentions, _ := api.GetStatusesMentionsTimeline(v)
	for _, tweet := range mentions {
		fmt.Println(tweet.Text)
	}
}
