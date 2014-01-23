package foosrank

import (
    "fmt"
    "code.google.com/p/gosqlite/sqlite"
)

func getDatabaseConnection() *sqlite.Conn {
    c, _ := sqlite.Open("games.db")
    return c
}

func mapGames(mapper func(Game), connection *sqlite.Conn) {
    var game Game
    var wId int
    var lId int
    query, err := connection.Prepare("SELECT winnerId, loserId, WinnerScore, LoserScore, GameId FROM Games;")
    err = query.Exec()
    if (err == nil) {
        for ; query.Next() ; {
            err := query.Scan(&wId, &lId, &game.WinnerScore, &game.LoserScore, &game.GameId)
            if err != nil { fmt.Println("unable to scan game") }
            game.Winner = getPlayerForId(wId, connection)
            game.Loser = getPlayerForId(lId, connection)
            mapper(game)
        }
    } else {
        fmt.Println("Unable to execute select games query")
    }
}

func getPlayerForId(id int, connection *sqlite.Conn) Player {
    var res Player
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


func addGameToDb(game Game, connection *sqlite.Conn) {
    fmt.Printf("Adding Game: %v\n", game)
    winnerId := getPlayerDbId(game.Winner, connection)
    loserId := getPlayerDbId(game.Loser, connection)
    sql := fmt.Sprintf("INSERT INTO Games(id, winnerId, loserId, WinnerScore, LoserScore, GameId) VALUES(NULL, %v, %v, %v, %v, %v);", winnerId, loserId, game.WinnerScore, game.LoserScore, game.GameId)
    err := connection.Exec(sql)
    if (err != nil) {
        fmt.Printf("Error adding %v: [%v]\n", game, err)
    }
}


func getPlayerDbId(player Player, connection *sqlite.Conn) int {
   id := getExistingPlayerDbId(player, connection)
   if (id > 0) {
        return id
    } else {
        addPlayerToDb(player, connection)
        id := getExistingPlayerDbId(player, connection)
        return id
    }
}

func getExistingPlayerDbId(player Player, connection *sqlite.Conn) int {
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


func addPlayerToDb(player Player, connection *sqlite.Conn) {
    fmt.Printf("Adding %v to Players table\n", player)
    sql := fmt.Sprintf("INSERT INTO Players(id, FirstName, LastName, PlayerId) VALUES(NULL, '%v', '%v', '%v');", player.FirstName, player.LastName, player.PlayerId)
    err := connection.Exec(sql)
    if (err != nil) {
        fmt.Printf("Error adding %v: [%v]\n", player, err)
    }
}
