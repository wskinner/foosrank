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
	GameId int64
}

type RankedPlayer struct {
    Player Player
    PlayerRank EloRank
}


type EloRank struct {
	Value float64
}

type RankingFunction func(float64, float64) (float64, float64)
