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
    PlayerRank EloRank
}

type EloRank struct {
	Rank int
}

type Rank struct {
    Mean float64
    StdDev float64
}

// This should be serialized and sent to the client
type Leaderboard struct {
	Players []RankedPlayer
}
