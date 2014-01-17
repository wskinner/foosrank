package foosrank

import (
	"github.com/wskinner/anaconda"
	"time"
	"fmt"
	"net/url"
)

func GetApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("WyWJsRBoWWQvnG8jnoO0bA")
	anaconda.SetConsumerSecret("1QxBiTHqcy5jeMgFzxwJyon5tdThphRzzOlmxiW3qG8")
	var api *anaconda.TwitterApi = anaconda.NewTwitterApi("2295398977-RCw9S9GM5UZ2lkQC3prlNfAKWgnaUdoRdcBEh6k", "olpgzfjU4a47f67kJQh7TrZ5ZjvOVE1RLUaD0RpbyD4JH")

	return api
}

func pollTwitter(api *anaconda.TwitterApi) []anaconda.Tweet {
	fmt.Println("polling twitter")
	v := url.Values{}

	// should be more than enough
	v.Set("count", "50")
	mentions, _ := api.GetStatusesMentionsTimeline(v)
	return mentions
}

func PollAtInterval(api *anaconda.TwitterApi, sleepTime time.Duration, tweetQueue chan anaconda.Tweet) {
	for {
		mentions := pollTwitter(api)
		fmt.Println("got ", len(mentions), " mentions")
		for _,t := range mentions {
			fmt.Println(t.Text)
			tweetQueue <- t
		}
		fmt.Println("sleeping")
		time.Sleep(sleepTime)
	}
}

