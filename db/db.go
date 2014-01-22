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
/*
func GetPlayerDbId(player foosrank.Player, connection *sqlite3.Conn) int {
   id := getExistingPlayerDbId(player, connection)
   if (id > 0) {
        return id
    } else {
        addPlayer(player, connection)
        id := getExistingPlayerDbId(player, connection)
        return id
    }
}
*/

func GetExistingPlayerDbId(player foosrank.Player, connection *sqlite.Conn) int {
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


func AddPlayer(player foosrank.Player, connection *sqlite.Conn) {
    fmt.Printf("Adding %v to Players table\n", player)
    sql := fmt.Sprintf("INSERT INTO Players(id, FirstName, LastName, PlayerId) VALUES(NULL, '%v', '%v', '%v');", player.FirstName, player.LastName, player.PlayerId)
    err := connection.Exec(sql)
    if (err != nil) {
        fmt.Printf("Error adding %v: [%v]\n", player, err)
    }
}
