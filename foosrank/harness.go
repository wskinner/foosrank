package foosrank

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "sort"
)

//an ordered list of RankedPlayers
type rankedPlayerSlice []*RankedPlayer
var leaderboard = make(rankedPlayerSlice, 0)
func (x rankedPlayerSlice) Len() int { return len(x) }
func (x rankedPlayerSlice) Less(i, j int) bool { return (x[i].PlayerRank.Value > x[j].PlayerRank.Value) }
func (x rankedPlayerSlice) Swap(i, j int) {
    var temp = x[i]
    x[i] = x[j]
    x[j] = temp
}


//a map of Person to RankedPlayer list Element
var players = make(map[Player]*RankedPlayer)

func readGameFile(rankingFunc RankingFunction) {
    file, err := ioutil.ReadFile("games.json")
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    var games []Game = make([]Game, 10)
    json.Unmarshal(file, &games)
    for _, game := range games {
        updateGame(&game, rankingFunc)
    }
    fmt.Println(leaderboard)
    sort.Sort(leaderboard)
}

func updateGame(game *Game, rankingFunc RankingFunction) {
    winner := game.Winner
    loser := game.Loser
    rankedWinner := addPlayer(winner, players, &leaderboard)
    rankedLoser := addPlayer(loser, players, &leaderboard)
    fmt.Println("before re-ranking: ", *rankedWinner, *rankedLoser)
    winnerNewRank, loserNewRank := rankingFunc(rankedWinner.PlayerRank.Value, rankedLoser.PlayerRank.Value)
    rankedWinner.PlayerRank.Value = winnerNewRank
    rankedLoser.PlayerRank.Value = loserNewRank
    fmt.Println("after re-ranking: ", *rankedWinner, *rankedLoser)
}


func addPlayer(p Player, ps map[Player]*RankedPlayer, leaders *rankedPlayerSlice) *RankedPlayer{
    if (ps[p] != nil) {
        fmt.Println("player: ", p, " already exists")
        return ps[p]
    } else {
        var rank = EloRank{1} //1 is default rank I guess
        var rankedPlayer = RankedPlayer{p, rank} //construct ranked player
        ps[p] = &rankedPlayer
        *leaders = append(*leaders, &rankedPlayer)
        fmt.Println("added player: ", p)
        return ps[p]
    }
}

//will ultimately output to a chan, just dont know what type yet
//will also take as arg a function, the ranking function.  Should conform to 
//some set interface so multiple ranking functions can be used
func ReadGames (gamesChan chan Game, rankingFunc RankingFunction) {
    readGameFile(rankingFunc)
    //for game := range gamesChan {
       //log game to master record
       //updateGame(&game, rankingFunc)
       //sort
       //publish leaderboard
    //}
}
