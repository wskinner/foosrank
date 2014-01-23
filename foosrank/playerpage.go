package foosrank
import (
	"encoding/json"
)

type Opponent struct {
	Player Player
	WinsAgainst int
	LossesAgainst int
	MyTotalPoints int
	TheirTotalPoints int
}

type PlayerData struct {
	Player RankedPlayer
	Opponents []Opponent
}

// Retrieve the appropriate data from the db, then serialize and send it to the connection
func updatePlayerPage(conn *connection, uid string) {
	//p1 := Player{"Will", "Skinner", "willskinner"}
	p2 := Player{"Michael", "Schiff", "michaelschiff"}
	data := []Opponent{Opponent{Player: p2}}
	bytes, _ := json.Marshal(data)
	conn.send <- bytes
}
