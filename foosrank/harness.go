package foosrank

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "container/list"
)



var leaderBoard = list.New()
var players = make(map[RankedPlayer]*list.Element)

func readGameFile() {
    file, err := ioutil.ReadFile("games.json")
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    var games []Game = make([]Game, 10)
    json.Unmarshal(file, &games)
    for _, game := range games {
        var winner = Person{game.Winner.FirstName, game.Winner.LastName}
        var loser = Person{game.Loser.FirstName, game.Loser.LastName}
        fmt.Println(winner, loser)

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
