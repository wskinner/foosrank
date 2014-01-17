package main

import (
	"github.com/wskinner/foosrank/foosrank"
	"fmt"
)

func main() {
	groups, err := foosrank.GetTweetEntities("will skinner 6 michael schiff 8")
	winner, loser, err := foosrank.GetPlayers(groups)

	fmt.Printf("winner: %+v\nloser: %+v\n", winner, loser)
	fmt.Printf("err: %v\n", err)

/*
	groups, err = foosrank.GetTweetEntities("will 6 michael schiff 8")
	winner, loser, err = foosrank.GetPlayers(groups)
	
	fmt.Println(winner, loser)
	fmt.Printf("err: %v\n", err)


	groups, err = foosrank.GetTweetEntities("will skinner 6 michael 8")
	winner, loser, err = foosrank.GetPlayers(groups)

	fmt.Println(winner, loser)
	fmt.Printf("err: %v\n", err)


	groups, err = foosrank.GetTweetEntities("will 6 michael 8")
	winner, loser, err = foosrank.GetPlayers(groups)

	fmt.Println(winner, loser)
	fmt.Printf("err: %v\n", err)
*/
}
