package foosrank

type Player struct {
	FirstName string
	LastName string
    PlayerId string
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
	Value int
}

type RankingFunction func(int, int) (int, int)
