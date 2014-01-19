package main
import (
	"github.com/wskinner/anaconda"
	"github.com/wskinner/foosrank/foosrank"
	"time"
	"fmt"
	"runtime"
)


func main() {
	fmt.Println("starting")
	tweetChan := make(chan anaconda.Tweet, 10)
	parsedChan := make(chan foosrank.Game, 10)
    leaderboardChan := make(chan []foosrank.RankedPlayer, 10)

	// 30 second intervals
	dur, err := time.ParseDuration("30000ms")
	if err != nil {
		fmt.Printf("Error: %v\n")
	}

	go foosrank.RunServer(leaderboardChan)

	go foosrank.PollAtInterval(foosrank.GetApi(), dur, tweetChan)
	go foosrank.ParseTweets(tweetChan, parsedChan)
    go foosrank.RankGames(parsedChan, foosrank.RankElo, leaderboardChan)
	
    // infinite loop
	for {
		runtime.Gosched()
	}
}
