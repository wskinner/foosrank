package foosrank

import (
	"encoding/json"
	"fmt"
	"time"
)

type hub struct {
	// Registered connections
	connections map[*connection]bool

	// Inbound messages from the connections
	broadcast chan []byte

	// Register requests from the connections
	register chan *connection

	// Unregister requests from the connections
	unregister chan *connection

	// The current leaderboard
	currentLeaderboard []RankedPlayer
}

var h = hub {
	broadcast: make(chan [] byte),
	register: make(chan *connection),
	unregister: make(chan *connection),
	connections: make(map[*connection]bool),
}

type Ping struct {
	Ping string
}

func ping() ([]byte, error) {
	pingMsg := Ping{Ping: "true"}
	return json.Marshal(pingMsg)
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			fmt.Println("New client connected.")
			h.connections[c] = true
			msg, _ := json.Marshal(h.currentLeaderboard)
			c.send <- msg
			// setup ping 
			go func () {
				for {
					pMsg, _ := ping()
					c.send <- pMsg
					t, _ := time.ParseDuration("280s")
					fmt.Println("Sent ping to client")
					time.Sleep(t)
				}
			}()
		case c := <-h.unregister:
			fmt.Println("Client disconnected.")
			delete(h.connections, c)
		case m := <-h.broadcast:
			// For each connection, send it the message. If the channel is not 
			// full, close the connection.
			fmt.Println("Broadcasting message to all clients: " + string(m))
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
					go c.ws.Close()
				}
			}
		}
	}
}


