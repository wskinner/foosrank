package foosrank

import (
	"fmt"
	"github.com/wskinner/anaconda"
	"regexp"
	"strconv"
	"errors"
	"strings"
)

// Matches and captures the form:
// fname [lname] score fname [lname] score
func GetTweetEntities(tweetstr string) ([]string, error) {
	matchstr := "([@a-zA-Z]+)\\s([@a-zA-Z]+\\s)?([0-9]+)\\s([@a-zA-Z]+)\\s([@a-zA-Z]+\\s)?([0-9]+)"
	matcher, err := regexp.Compile(matchstr)
	if err != nil {
		fmt.Printf("Error compile regular expression: %v\n", err)
		return nil, err
	}

	matched := matcher.FindStringSubmatch(tweetstr)
	
	// Should have at least 2 first names and 2 scores
	// 5 because the first index is occupied by entire string
	if len(matched) < 5 {
		err = errors.New("Not enough game parameters")
		fmt.Printf("%v\n", err)
		return nil, err
	}

	return matched[1:], err
}

func updateString(p *Player, val string) error {
	var err error = nil
	if (*p).FirstName == "" {
		(*p).FirstName = val
	} else if (*p).LastName == "" {
		(*p).LastName = val
	} else {
		err = errors.New("Error: Too many string values in game")
	}
	return err
}

// Return (winner, loser, winnerscore, loserscore, error)
func GetPlayers(groups []string) (Player, Player, int, int, error) {
	var err error = nil
	var p1 Player
	var p2 Player
	var s1 int
	var s2 int

	p1ScoreSet := false
	for _, s := range groups {
		fmt.Println(s)
		// it's a @mention, so skip it
		if strings.HasPrefix(s, "@") {
			continue
		}
		i, e := strconv.ParseInt(s, 10, 0)
		// it's an int
		if e == nil {
			if p1ScoreSet {
				s2 = int(i)
			} else {
				s1 = int(i)
				p1ScoreSet = true
			}
		// it's a string
		} else {
			if p1ScoreSet {
				err = updateString(&p2, s)
			} else {
				err = updateString(&p1, s)
			}
		}
	}

	p1.PlayerId = p1.FirstName+p1.LastName
    p2.PlayerId = p2.FirstName+p2.LastName
    if s1 > s2 {
		return p1, p2, s1, s2, err
	} else {
		return p2, p1, s2, s1, err
	}


}

func parseTweet(tweet anaconda.Tweet) (Game, error) {
	var game Game
	groups, err := GetTweetEntities(tweet.Text)
	if err != nil {
		return game, err
	}

	game.Winner, game.Loser, game.WinnerScore, game.LoserScore, err = GetPlayers(groups)

	fmt.Printf("parsed a game:\n%+v\n", game)
	
	return game, err
}


// Block on reading from tweetChan, and write parsed Game structs
// to parsedChan in real time.
func ParseTweets (tweetChan chan anaconda.Tweet, parsedChan chan Game) {
	for t := range tweetChan {
		fmt.Println("parsing tweet with text: ", t.Text)
		game, err := parseTweet(t)
		if err == nil {
			parsedChan <- game
		} else {
			fmt.Printf("%v\n", err)
		}
	}
}
