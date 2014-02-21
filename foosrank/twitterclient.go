package foosrank

import (
	"github.com/wskinner/anaconda"
	"time"
	"fmt"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"math"
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
	fmt.Println("polling twitter - lastid=", getLastId())
	v := url.Values{"since_id": {getLastId()}}
	//v := url.Values{}

	// should be more than enough
	v.Set("count", "199")
	mentions, _ := api.GetStatusesMentionsTimeline(v)
	sort.Sort(anaconda.ById(mentions))
	if len(mentions) == 1 {
		saveLastId(strconv.FormatInt(mentions[0].Id, 10))
	} else if len(mentions) > 1 {
		saveLastId(strconv.FormatInt(mentions[len(mentions)-1].Id, 10))
	}

	return mentions
}

func saveLastId(id string) {
	bytes := []byte(id)
	err := ioutil.WriteFile("lastid.txt", bytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// If no last id, just get everything
func getLastId() string {
	dat, err := ioutil.ReadFile("lastid.txt")
	if err != nil {
		max := strconv.FormatUint(math.MaxUint64, 10)
		saveLastId(max)
		return max
	}
	return string(dat)
}

func PollAtInterval(api *anaconda.TwitterApi, sleepTime time.Duration, tweetQueue chan anaconda.Tweet) {
	fmt.Println("in PollAtInterval")
	for {
		mentions := pollTwitter(api)
		fmt.Println("got ", len(mentions), " mentions")
		for _,t := range mentions {
			fmt.Printf("Text: %s. Id: %d\n", t.Text, t.Id)
			tweetQueue <- t
		}
		fmt.Println("sleeping")
		time.Sleep(sleepTime)
	}
}

