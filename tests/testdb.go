package main

import (
    "fmt"
    "github.com/wskinner/foosrank/foosrank"
)

func f(g foosrank.Game) {
    fmt.Println(g)
}

func main() {
    connection := GetDatabaseConnection()
    player := foosrank.Player{"Michael", "Schiff", "michaelschiff"}
    other := foosrank.Player{"Will", "Skinner", "willskinner"}
    fmt.Println(GetPlayerDbId(player, connection))
    game := foosrank.Game{player, other, 8, 6, 0}
    AddGame(game, connection)
    fmt.Println(GetPlayerForId(1, connection))
    MapGames(f, connection)
}
