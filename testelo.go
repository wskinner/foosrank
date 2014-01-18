package main

import (
	"fmt"
	"github.com/wskinner/foosrank/foosrank"
)


func main() {
	score1 := 1500
	score2 := 1500
	for i := 0; i < 5; i++ {
		score1,score2 = foosrank.RankElo(score1, score2)
	}
	fmt.Printf("Winner's old score: %d\nLoser's old score: %d\nWinner's new score: %d\nLoser's new score: %d\n", 1500, 1500, score1, score2)
}


