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

func getExistingUidDbId(uid string, connection *sqlite.Conn) (id int, err error) {
	sql := fmt.Sprintf("SELECT id FROM Players WHERE Players.PlayerId = '%v';", uid)
	fmt.Printf("Executing query: %v\n", sql)
	query, err := connection.Prepare(sql)
	err = query.Exec()
	if err == nil && query.Next() {
		query.Scan(&id)
		return id, err
	} else {
		fmt.Printf("Error: %v\n", err)
		return -1, err
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

func getAllOpponents(uid string, connection *sqlite.Conn) (opponents []Opponent) {
	id, _ := getExistingUidDbId(uid, connection)
	results := make(map[string]*Opponent)
	wins := fmt.Sprintf(`
			SELECT FirstName, LastName, SUM(WinnerScore) as MyTotalPoints, SUM(LoserScore) as TheirTotalPoints, COUNT(*) as WinsAgainst, Players.PlayerId as oppId
			FROM Games, Players
			WHERE Games.loserId=Players.id AND Games.winnerId=%v
			GROUP BY oppId; 
			`, id)
	losses := fmt.Sprintf(`
			SELECT FirstName, LastName, SUM(LoserScore) as MyTotalPoints, SUM(WinnerScore) as TheirTotalPoints, COUNT(*) as LossesAgainst, Players.PlayerId as oppId
			FROM Games, Players
			WHERE Games.winnerId=Players.id AND Games.loserId=%v
			GROUP BY oppId
			`, id)
	fmt.Printf("Executing query:\n%s\n", wins)
	query, _ := connection.Prepare(wins)
	err := query.Exec()
	if err != nil {
		fmt.Println(err)
	}
        for ; query.Next() ; {
        	player := Player{}
        	opp := Opponent{}
		err = query.Scan(&player.FirstName, &player.LastName, &opp.MyTotalPoints, &opp.TheirTotalPoints, &opp.WinsAgainst, &player.PlayerId)
		if err != nil { 
			fmt.Printf("Error: %v\n", err)
			//continue
		}
		opp.Player = player
		fmt.Println(player)
		fmt.Println(opp)
		results[player.PlayerId] = &opp
        }

	fmt.Printf("Executing query:\n%s\n", losses)
	query, _ = connection.Prepare(losses)
	err = query.Exec()
	if err != nil {
		fmt.Println(err)
	}
        for ; query.Next() ; {
        	player := Player{}
        	opp := Opponent{}
		err = query.Scan(&player.FirstName, &player.LastName, &opp.MyTotalPoints, &opp.TheirTotalPoints, &opp.LossesAgainst, &player.PlayerId)
		if err != nil { 
			fmt.Printf("Error: %v\n", err)
			//continue
		}
		opp.Player = player
		fmt.Println(player)
		fmt.Println(opp)
		if val,ok := results[player.PlayerId]; !ok {
			results[player.PlayerId] = &opp
		} else {
			merge(val, &opp)
		}
        }
        for _,v := range results {
        	opponents = append(opponents, *v)
        }
        return opponents
}

// Merge opp into target
func merge(target *Opponent, opp *Opponent) {
	target.WinsAgainst += opp.WinsAgainst
	target.LossesAgainst += opp.LossesAgainst
	target.MyTotalPoints += opp.MyTotalPoints
	target.TheirTotalPoints += opp.TheirTotalPoints
}
