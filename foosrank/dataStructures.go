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

