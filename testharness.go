package main

import (
    "github.com/wskinner/foosrank/foosrank"
)

func main() {
    foosrank.ReadGames(make(chan foosrank.Game))
}

