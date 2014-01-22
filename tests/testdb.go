package main

import (
    "fmt"
    "github.com/wskinner/foosrank/foosrank"
    "github.com/wskinner/foosrank/db"
)

func main() {
    connection := db.GetDatabaseConnection()
    player := foosrank.Player{"Michael", "Schiff", "michaelschiff"}
    fmt.Println(db.GetPlayerDbId(player, connection))
}
