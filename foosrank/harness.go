// +build ignore

package foosrank

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "container/list"
)

var leaderBoard = list.New()
var players = make(map[RankedPlayer]bool)

func readGameFile() {
    file, err := ioutil.ReadFile("games.json")
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    var games []Game
    json.Unmarshal(file, games)
    fmt.Println(games)
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
