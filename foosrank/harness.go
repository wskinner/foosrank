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

var (
    //a map of Player (soon a unique id field in a player) to *RankedPlayer in leaderboard
    players = make(map[string]*RankedPlayer)

    //a sorted slice of RankedPlayers
    leaderboard = make(rankedPlayerSlice, 0)
)

//loads games out of file, updates the leaderboard for each game, 
//then resorts leaderboard after all updates
func readGameFile(rankingFunc RankingFunction) {
    file, err := ioutil.ReadFile("games.json")
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    var games []Game = make([]Game, 20)
    json.Unmarshal(file, &games)
    for _, game := range games {
    	fmt.Println("Game: ", game)
        updateGame(game, rankingFunc)
    }
    fmt.Println(leaderboard)
}

//func getRankingMapFunction(rankingFunc RankingFunction) {
//    return func(game Game) {
//        updateGame(

//gets the pointer to RankedPlayer from players map (via addPlayer)
//finds new rank for players, and updates the RankedPlayer structs in the leaderboard
func updateGame(game Game, rankingFunc RankingFunction) {
    winner := game.Winner
    loser := game.Loser
    rankedWinner := addPlayer(winner, players, &leaderboard)
    rankedLoser := addPlayer(loser, players, &leaderboard)
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
func addPlayer(p Player, ps map[string]*RankedPlayer, leaders *rankedPlayerSlice) *RankedPlayer{
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

func getLogFile() *os.File {
    file, err := os.OpenFile("games.json", os.O_RDWR | os.O_APPEND, 0666)
    if err == nil {
        return file
    } else {
        fmt.Println("error: couldn't open log")
        return nil
    }
}

func logGame(game Game, file *os.File) {
    fi, _ := file.Stat()
    bytes, err := json.Marshal(game)
    if err == nil {
        sep := "\n"
        if (fi.Size() > 3) { sep = ","+sep }
        n, e := file.WriteAt([]byte(sep), fi.Size()-2)
        if e != nil {
        	fmt.Println("Error: ", e)
        } else {
        	fmt.Printf("Wrote %d bytes\n", n)
        }
        file.Write(bytes)
        file.Write([]byte("\n]"))
        fmt.Println("logged game")
    } else {
        fmt.Println("failed to log game")
    }
}

//will ultimately output to a chan, just dont know what type yet
//will also take as arg a function, the ranking function.  Should conform to 
//some set interface so multiple ranking functions can be used
func RankGames (gamesChan chan Game, rankingFunc RankingFunction, leaderboardChan chan []RankedPlayer) {
    readGameFile(rankingFunc)
    sort.Sort(leaderboard)
    leaderboardChan <- convertToValues(leaderboard)

    gameLog := getLogFile()

    for game := range gamesChan {
       logGame(game, gameLog)
       updateGame(game, rankingFunc)
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
