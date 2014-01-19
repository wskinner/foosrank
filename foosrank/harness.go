package foosrank

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "container/list"
)



//an ordered list of RankedPlayers
var leaderBoard = list.New()

//a map of Person to RankedPlayer list Element
var players = make(map[Player]bool)

func readGameFile() {
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
        fmt.Println(winner, loser)
        addPlayer(winner, players)
        addPlayer(loser, players)
    }
}

func addPlayer(p Player, ps map[Player]bool) bool {
    if (ps[p]) {
        fmt.Println("player: ", p, " already exists")
        return false
    } else {
        ps[p] = true
        fmt.Println("added player: ", p)
        return true
    }
}

//will ultimately output to a chan, just dont know what type yet
//will also take as arg a function, the ranking function.  Should conform to 
//some set interface so multiple ranking functions can be used
func ReadGames (gamesChan chan Game, rankingFunc RankingFunction) {
    readGameFile()
    //for game := range gamesChan {
       //log game to master record
       //add ranks to game players
       //ask ranking function for updated ranks
       //update and publish leaderboard
    //}
}
