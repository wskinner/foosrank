package foosrank

import (
	"github.com/wskinner/anaconda"
	"time"
	"fmt"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"os"
)

type TwitterCreds struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

func GetApi() *anaconda.TwitterApi {
	file, e := ioutil.ReadFile("credentials.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var creds TwitterCreds
	json.Unmarshal(file, &creds)
	fmt.Println(creds)
	anaconda.SetConsumerKey(creds.ConsumerKey)
	anaconda.SetConsumerSecret(creds.ConsumerSecret)
	var api *anaconda.TwitterApi = anaconda.NewTwitterApi(creds.AccessToken, creds.AccessTokenSecret)

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

