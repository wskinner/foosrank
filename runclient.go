package main
import (
	"github.com/wskinner/anaconda"
	"github.com/wskinner/foosrank"
	"time"
	"fmt"
	"runtime"
)


func main() {
	fmt.Println("starting")
	tweetChan := make(chan anaconda.Tweet, 10)
	dur, err := time.ParseDuration("5000ms")
	if err == nil {
		fmt.Println("error was nil")
		go foosrank.PollAtInterval(foosrank.GetApi(), dur, tweetChan)
		for {
			runtime.Gosched()
		}
	}
}
