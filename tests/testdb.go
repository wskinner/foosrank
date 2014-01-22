package main

import (
    "github.com/wskinner/foosrank/foosrank"
    "github.com/wskinner/foosrank/db"
)

func main() {
    connection := db.GetDatabaseConnection()
    player := foosrank.Player{"Michael", "Schiff", "michaelschiff"}
    db.AddPlayer(player, connection)
}
