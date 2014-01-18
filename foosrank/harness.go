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
        
        if (players[winner]) {
            fmt.Println("skipping", winner)
        } else {
            players[winner] = true
            fmt.Println("adding", winner)
        }
        
        if (players[loser]) {
            fmt.Println("skipping", loser)
        } else {
            fmt.Println("adding", loser)
            players[loser] = true
        }
    }
}

//will ultimately output to a chan, just dont know what type yet
func ReadGames (gamesChan chan Game) {
    readGameFile()
    //for game := range gamesChan {
       //log game to master record
       //add ranks to game players
       //ask ranking function for updated ranks
       //update and publish leaderboard
    //}
}
