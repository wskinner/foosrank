package db

import (
    "fmt"
    "code.google.com/p/gosqlite/sqlite"
    "github.com/wskinner/foosrank/foosrank"
)

func GetDatabaseConnection() *sqlite.Conn {
    c, _ := sqlite.Open("games.db")
    return c
}

func MapGames(mapper func(foosrank.Game), connection *sqlite.Conn) {
    var game foosrank.Game
    var wId int
    var lId int
    query, err := connection.Prepare("SELECT winnerId, loserId, WinnerScore, LoserScore, GameId FROM Games;")
    err = query.Exec()
    if (err == nil) {
        for ; query.Next() ; {
            err := query.Scan(&wId, &lId, &game.WinnerScore, &game.LoserScore, &game.GameId)
            if err != nil { fmt.Println("unable to scan game") }
            game.Winner = GetPlayerForId(wId, connection)
            game.Loser = GetPlayerForId(lId, connection)
            mapper(game)
        }
    } else {
        fmt.Println("Unable to execute select games query")
    }
}

func GetPlayerForId(id int, connection *sqlite.Conn) foosrank.Player {
    var res foosrank.Player
    sql := fmt.Sprintf("SELECT Players.FirstName, Players.LastName, Players.PlayerId FROM Players WHERE Players.id = %v", id)
    query, err := connection.Prepare(sql)
    err = query.Exec()
    if (err == nil && query.Next()) {
        err = query.Scan(&res.FirstName, &res.LastName, &res.PlayerId)
        if (err == nil) {
            return res
        } else {
            fmt.Printf("error scanning player for id %v : %v\n", id, err)
            return res
        }
    } else {
        fmt.Printf("error executing query for player %v : %v\n", id, err)
        return res
    }
}


func AddGame(game foosrank.Game, connection *sqlite.Conn) {
    fmt.Printf("Adding Game: %v\n", game)
    winnerId := GetPlayerDbId(game.Winner, connection)
    loserId := GetPlayerDbId(game.Loser, connection)
    sql := fmt.Sprintf("INSERT INTO Games(id, winnerId, loserId, WinnerScore, LoserScore, GameId) VALUES(NULL, %v, %v, %v, %v, %v);", winnerId, loserId, game.WinnerScore, game.LoserScore, game.GameId)
    err := connection.Exec(sql)
    if (err != nil) {
        fmt.Printf("Error adding %v: [%v]\n", game, err)
    }
}


func GetPlayerDbId(player foosrank.Player, connection *sqlite.Conn) int {
   id := getExistingPlayerDbId(player, connection)
   if (id > 0) {
        return id
    } else {
        addPlayer(player, connection)
        id := getExistingPlayerDbId(player, connection)
        return id
    }
}

func getExistingPlayerDbId(player foosrank.Player, connection *sqlite.Conn) int {
    sql := fmt.Sprintf("SELECT id FROM Players WHERE Players.FirstName = '%v' AND Players.LastName = '%v' AND Players.PlayerId = '%v';", player.FirstName, player.LastName, player.PlayerId)
    query, err := connection.Prepare(sql)
    err = query.Exec()
    if (err == nil && query.Next()) {
        var id int
        query.Scan(&id)
        return id
    } else {
        return -1
    }
}


func addPlayer(player foosrank.Player, connection *sqlite.Conn) {
    fmt.Printf("Adding %v to Players table\n", player)
    sql := fmt.Sprintf("INSERT INTO Players(id, FirstName, LastName, PlayerId) VALUES(NULL, '%v', '%v', '%v');", player.FirstName, player.LastName, player.PlayerId)
    err := connection.Exec(sql)
    if (err != nil) {
        fmt.Printf("Error adding %v: [%v]\n", player, err)
    }
}
