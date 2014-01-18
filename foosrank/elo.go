package foosrank

import "math"

// Return (winnerNewRank, loserNewRank)
func RankElo(winnerRank int, loserRank int) (int, int) {
	rankDiff := winnerRank - loserRank
	exp := float64(-1 * rankDiff / 400)
	odds := 1/ (1 + math.Pow(10, exp))
	var k int
	if winnerRank < 2100 {
		k = 32
	} else if winnerRank >= 2100 && winnerRank < 2400 {
		k = 24
	} else {
		k = 16
	}
	
	winnerNewRank := int(float64(winnerRank) + float64(k) * (1 - odds))
	loserNewRank := loserRank - (winnerNewRank - winnerRank)
	return winnerNewRank, loserNewRank
}

