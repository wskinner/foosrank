package foosrank

import (
	"math"
	"fmt"
)

// Return (winnerNewRank, loserNewRank)
func RankElo(winnerRank float64, loserRank float64) (float64, float64) {
	rankDiff := float64(winnerRank - loserRank)
	exp := float64(-1.0 * rankDiff / 400.0)
	odds := 1.0 / (1.0 + math.Pow(10, exp))
	fmt.Println("Odds of winning: ", odds)
	var k float64
	if winnerRank < 2100 {
		k = 32
	} else if winnerRank >= 2100 && winnerRank < 2400 {
		k = 24
	} else {
		k = 16
	}
	
	winnerNewRank := winnerRank + k * (1.0 - odds)
	loserNewRank := loserRank - (winnerNewRank - winnerRank)
	return winnerNewRank, loserNewRank
}

