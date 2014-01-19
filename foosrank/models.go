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
    PlayerId int64 // who knows, maybe we will have billions of players
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
