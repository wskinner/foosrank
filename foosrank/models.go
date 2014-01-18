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

type Person struct {
    FirstName string
    LastName string
}

type RankedPlayer struct {
    Player Person
    PlayerRank Rank
}

type EloRank struct {
	Rank int
}

type Rank struct {
    Mean float64
    StdDev float64
}
