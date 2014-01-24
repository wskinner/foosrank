package foosrank
import (
	"encoding/json"
	"code.google.com/p/gosqlite/sqlite"
)

type Opponent struct {
	Player Player
	WinsAgainst int
	LossesAgainst int
	MyTotalPoints int
	TheirTotalPoints int
}

// Retrieve the appropriate data from the db, then serialize and send it to the connection
func updatePlayerPage(conn *connection, uid string, dbConn *sqlite.Conn) {
	data := getAllOpponents(uid, dbConn)
	bytes, _ := json.Marshal(data)
	conn.send <- bytes
}
