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
    //without just first/last name coming from the parser, its impossible to uniquely id a player
    //either there is no one with that name, so give them the next available id
    //or: someone already has that name, and we cant tell if this is a new person with
    //the same name, or someone we've already seen.  Names should be a unique player id instead
    //(supplied directly in the tweet).  Players should have to register this name with the site
    //before using it.
    //PlayerId int64 // who knows, maybe we will have billions of players
}

type Rank interface {
    Rank() int
}

type EloRank struct {
	Value int
}
func (eloRank EloRank) Rank() int {
    return eloRank.Value
}

type TrueSkillRank struct {
    Mean int
    StdDev int
}
func (trueSkill TrueSkillRank) Rank() int {
    return trueSkill.Mean
}

type RankingFunction func(int, int) (int, int)

// This should be serialized and sent to the client
type Leaderboard struct {
	Players []RankedPlayer
}

