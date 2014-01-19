package main
import (
	"github.com/wskinner/anaconda"
	"github.com/wskinner/foosrank/foosrank"
	"time"
	"fmt"
	"runtime"
    "os"
)


func main() {
	fmt.Println("starting")
	tweetChan := make(chan anaconda.Tweet, 10)
	parsedChan := make(chan foosrank.Game, 10)
	leaderboardChan := make(chan []foosrank.RankedPlayer, 10)

	go foosrank.RunServer(leaderboardChan)

	if len(os.Args) < 2 {
        fmt.Println("specify a source of data. either 'twitter' or 'cmdline'")
        os.Exit(1)
    }
    if (os.Args[1] == "twitter") {
        //10 second intervals
        dur, err := time.ParseDuration("10000ms")
        if err != nil {
            fmt.Printf("Error: %v\n")
        }
        go foosrank.PollAtInterval(foosrank.GetApi(), dur, tweetChan)
    } else if (os.Args[1] == "cmdline") {
        go foosrank.ReadTweetFromStdIn(tweetChan)
    } else {
        fmt.Println("unknown argument, use either 'twitter' or 'cmdline'")
    }

	go foosrank.ParseTweets(tweetChan, parsedChan)

	go foosrank.RankGames(parsedChan, foosrank.RankElo, leaderboardChan)
	
    // infinite loop
	for {
		runtime.Gosched()
	}
}
