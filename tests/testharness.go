package main

import (
    "github.com/wskinner/foosrank/foosrank"
)

func main() {
    foosrank.RankGames(make(chan foosrank.Game), foosrank.RankElo, make(chan []foosrank.RankedPlayer))
}

