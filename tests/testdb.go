package main

import (
    "fmt"
    "github.com/wskinner/foosrank/foosrank"
    "github.com/wskinner/foosrank/db"
)

func main() {
    connection := db.GetDatabaseConnection()
    player := foosrank.Player{"Michael", "Schiff", "michaelschiff"}
    other := foosrank.Player{"Will", "Skinner", "willskinner"}
    fmt.Println(db.GetPlayerDbId(player, connection))
    game := foosrank.Game{player, other, 8, 6, 0}
    db.AddGame(game, connection)
}
