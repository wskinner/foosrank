package foosrank

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "container/list"
    //"sort"
)



//an ordered list of RankedPlayers
var leaderboard = list.New()

//a map of Person to RankedPlayer list Element
var players = make(map[Player]*list.Element)

func readGameFile(rankingFunc RankingFunction) {
    file, err := ioutil.ReadFile("games.json")
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    var games []Game = make([]Game, 10)
    json.Unmarshal(file, &games)
    for _, game := range games {
	    winner := game.Winner
	    loser := game.Loser
        rankedWinner := addPlayer(winner, players, leaderboard).Value.(RankedPlayer)
        rankedLoser := addPlayer(loser, players, leaderboard).Value.(RankedPlayer)
        fmt.Println(rankedWinner, rankedLoser)
        winnerNewRank, loserNewRank := rankingFunc(rankedWinner.PlayerRank.Value, rankedLoser.PlayerRank.Value)
        fmt.Println(winnerNewRank, loserNewRank)
    }
    //sort.Sort(leaderboard) //need to implement Sort interface for my list
}

//TODO: implement updateGame function, does the stuff above since is reusable

func addPlayer(p Player, ps map[Player]*list.Element, leaders *list.List) *list.Element{
    if (ps[p] != nil) {
        fmt.Println("player: ", p, " already exists")
        return ps[p]
    } else {
        var rank = EloRank{1} //1 is default rank I guess
        var rankedPlayer = RankedPlayer{p, rank} //construct ranked player
        ps[p] = leaders.PushBack(rankedPlayer) //map player to ranked player
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
       //add ranks to game players
       //ask ranking function for updated ranks
       //update and publish leaderboard
    //}
}
