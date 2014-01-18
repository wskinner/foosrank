package foosrank

type Player struct {
	FirstName string
	LastName string
}

type Game struct {
	Winner Player
	Loser Player
	WinnerScore int
	LoserScore int
}

type RankedPlayer struct {
    Player Player
    PlayerRank Rank
}

type EloRank struct {
	Rank int
}

type Rank struct {
    Mean float64
    StdDev float64
}
