package foosrank

type Player struct {
	FirstName string
	LastName string
	Score int
}

type Game struct {
	Winner Player
	Loser Player
}

type RankedPlayer struct {
    FirstName string
    LastName string
    PlayerRank Rank
}

type Rank struct {
    Mean float64
    StdDev float64
}
