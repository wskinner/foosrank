package foosrank

import (
    "fmt"
    "sort"
    "time"
    "code.google.com/p/gosqlite/sqlite"
)

type rankedPlayerSlice []*RankedPlayer
func (x rankedPlayerSlice) Len() int { return len(x) }
func (x rankedPlayerSlice) Less(i, j int) bool { return (x[i].PlayerRank.Value > x[j].PlayerRank.Value) }
func (x rankedPlayerSlice) Swap(i, j int) {
    var temp = x[i]
    x[i] = x[j]
    x[j] = temp
}

var (
    players = make(map[string]*RankedPlayer)
    leaderboard = make(rankedPlayerSlice, 0)
)

func getRankingMapFunction(rankingFunc RankingFunction) func(Game) {
    return func(game Game) {
        updateGame(game, rankingFunc)
    }
}

//gets the pointer to RankedPlayer from players map (via addPlayerToLeaderboard)
//finds new rank for players, and updates the RankedPlayer structs in the leaderboard
func updateGame(game Game, rankingFunc RankingFunction) {
    winner := game.Winner
    loser := game.Loser
    rankedWinner := addPlayerToLeaderboard(winner, players, &leaderboard)
    rankedLoser := addPlayerToLeaderboard(loser, players, &leaderboard)
    fmt.Println("before re-ranking: ", *rankedWinner, *rankedLoser)

    winnerOldRank := rankedWinner.PlayerRank.Value
    loserOldRank := rankedLoser.PlayerRank.Value

    winnerNewRank, loserNewRank := rankingFunc(winnerOldRank, loserOldRank)
    rankedWinner.PlayerRank.Value = winnerNewRank
    rankedLoser.PlayerRank.Value = loserNewRank

    rankedWinner.PlayerRankDelta = winnerNewRank - winnerOldRank
    rankedLoser.PlayerRankDelta = loserNewRank - loserOldRank
    fmt.Println("after re-ranking: ", *rankedWinner, *rankedLoser)
}


//checks to see if we have a mapping for this player in the players map
//if so, we return the *RankedPlayer it maps to
//if not, we create a new ranked player, associate the Player with * to new RankedPlayer
//and add *RankedPlayer to end of leaderboard
func addPlayerToLeaderboard(p Player, ps map[string]*RankedPlayer, leaders *rankedPlayerSlice) *RankedPlayer{
    var id = p.PlayerId
    if (ps[id] != nil) {
        fmt.Println("player: ", p, " already exists")
        return ps[id]
    } else {
        var rank = EloRank{1500} //1 is default rank I guess
        var rankedPlayer = RankedPlayer{p, rank, 0.0} //construct ranked player
        ps[id] = &rankedPlayer
        *leaders = append(*leaders, &rankedPlayer)
        fmt.Println("added player: ", p)
        return ps[id]
    }
}

func RankGames (gamesChan chan Game, rankingFunc RankingFunction, leaderboardChan chan []RankedPlayer) {
    var dbConnection = getDatabaseConnection()
    mapGames(getRankingMapFunction(rankingFunc), dbConnection)
    sort.Sort(leaderboard)
    leaderboardChan <- convertToValues(leaderboard)
    for game := range gamesChan {
    	addGameUpdateLeaders(game, rankingFunc, dbConnection, 0)
    	leaderboardChan <- convertToValues(leaderboard)
    }
}

func addGameUpdateLeaders(game Game, rankingFunc RankingFunction, connection *sqlite.Conn, retryTimes int) {
  err := addGameToDb(game, connection)
  if (err == nil) {
    updateGame(game, rankingFunc)
    sort.Sort(leaderboard)
  } else {
    if (retryTimes < 3) {
      dur, _ := time.ParseDuration("100ms")
      time.Sleep(dur)
      addGameUpdateLeaders(game, rankingFunc, connection, retryTimes+1)
    } else {
      fmt.Printf("Unable to add game: %v after 3 tries\n", game)
    }
  }
}

func convertToValues(x rankedPlayerSlice) []RankedPlayer {
    res := make([]RankedPlayer, len(x))
    for i := 0; i < len(x); i++ {
       res[i] = *(x[i])
    }
    return res
}
