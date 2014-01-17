package foosrank

import (
	"fmt"
	"github.com/wskinner/anaconda"
	"regexp"
	"strings"
	"strconv"
)

// Matches and captures the form:
// fname [lname] score fname [lname] score
func GetTweetEntities(tweetstr string) []string {
	matchstr := "([a-zA-Z]+)\\s([a-zA-Z]+\\s)?([0-9]+)\\s([a-zA-Z]+)\\s([a-zA-Z]+\\s)?([0-9]+)"
	matcher, err := regexp.Compile(matchstr)
	if err != nil {
		fmt.Printf("Error compile regular expression: %v\n", err)
	}

	matched := matcher.FindStringSubmatch(tweetstr)
	fmt.Println(matched[1:])
	return matched
}

func getPlayers(groups []string) (winner Player, loser Player) {
	var err Error
	var p1, p2, Player
	func updateString(p1 Player, p2 Player, val string) {
		if p1.FirstName == nil {
			p1.FirstName = val
		} else if p1.LastName == nil {
			p1.LastName = val
		} else if p2.FirstName == nil {
			p2.FirstName = val
		} else if p2.LastName == nil {
			p2.LastName = val
		} else {
			fmt.Printf("Error: Too many string values in game")
		}
	}

	func updateInt(p1 Player, p2 Playerl, val int) {
		if p1.Score == nil {
			p1.Score = val
		} else if p2 == nil {
			p2.Score = val
		} else {
			fmt.Printf("Error: Too many integer values in game")
		}
	}

	for s := groups {
		i, e := strconv.ParseInt(s)
		if e == nil {
			updateInt(p1, p2, i)
		} else {
			updateString(p1, p2, s)
		}
	}

	if 
}

func parseTweet(tweet anaconda.Tweet) Game {
	groups := getTweetEntities(tweet.Text)

	var game Game
	winner, loser := getPlayers(groups)
	game.Winner = winner
	game.Loser = loser
	
	return game
}

// Block on reading from tweetChan, and write parsed Game structs
// to parsedChan in real time.
func ParseTweets (tweetChan chan anaconda.Tweet, parsedChan chan Game) {
	for t := range tweetChan {
		parsedChan <- parseTweet(t)
	}
}
