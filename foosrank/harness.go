// +build ignore

package foosrank

import (
    "os"
)

func getLog() *os.File {
    file, err := os.Open("games.log")
    if err == nil {
        return file
    } else {
        return nil
    }
}

//will ultimately output to a chan, just dont know what type yet
func ReadGames (gamesChan chan Game) {
    logFile *os.File = getLog()
    for game := range gamesChan {
       //log game to master record
       //add ranks to game players
       //ask ranking function for updated ranks
       //update and publish leaderboard
    }
}
