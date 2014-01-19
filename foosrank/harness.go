package foosrank

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "sort"
)

type rankedPlayerSlice []*RankedPlayer
func (x rankedPlayerSlice) Len() int { return len(x) }
func (x rankedPlayerSlice) Less(i, j int) bool { return (x[i].PlayerRank.Value > x[j].PlayerRank.Value) }
func (x rankedPlayerSlice) Swap(i, j int) {
    var temp = x[i]
    x[i] = x[j]
    x[j] = temp
}

//a map of Player (soon a unique id field in a player) to *RankedPlayer in leaderboard
var players = make(map[Player]*RankedPlayer)

//a sorted slice of RankedPlayers
var leaderboard = make(rankedPlayerSlice, 0)

//loads games out of file, updates the leaderboard for each game, 
//then resorts leaderboard after all updates
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
}

//gets the pointer to RankedPlayer from players map (via addPlayer)
//finds new rank for players, and updates the RankedPlayer structs in the leaderboard
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


//checks to see if we have a mapping for this player in the players map
//if so, we return the *RankedPlayer it maps to
//if not, we create a new ranked player, associate the Player with * to new RankedPlayer
//and add *RankedPlayer to end of leaderboard
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
func RankGames (gamesChan chan Game, rankingFunc RankingFunction, leaderboardChan chan []RankedPlayer) {
    readGameFile(rankingFunc)
    sort.Sort(leaderboard)
    leaderboardChan <- convertToValues(leaderboard)
    for game := range gamesChan {
       //log game to master record
       updateGame(&game, rankingFunc)
       sort.Sort(leaderboard)
       leaderboardChan <- convertToValues(leaderboard)
    }
}

func convertToValues(x rankedPlayerSlice) []RankedPlayer {
    res := make([]RankedPlayer, len(x))
    for i := 0; i < len(x); i++ {
       res[i] = *(x[i])
    }
    return res
}
