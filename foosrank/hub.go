package foosrank

import (
	"encoding/json"
	"fmt"
	"time"
)

type PlayerPageRegistration struct {
	Uid string
	Conn *connection
}

// Websocket hub for leaderboards only
type hub struct {
	// Registered connections (home page only)
	connections map[*connection]bool

	// Broadcast leaderboards on this chan 
	broadcast chan []byte

	// Register requests from the connections
	register chan *connection

	// Unregister requests from the connections
	unregister chan *connection

	// The current leaderboard
	currentLeaderboard []RankedPlayer

	///////////////////////////////////////////////////////////
	// Unregister requests from player page connections
	pUnregister chan *connection

	// Registered connections (players page only)
	// uid -> *connection
	pConnections map[string]*connection

	pConnectionsSet map[*connection]string

	// Register connections (players page only)
	pRegister chan PlayerPageRegistration

	// Players who have updated information
	updatedUids chan string
}

var h = hub {
	broadcast: make(chan [] byte),
	register: make(chan *connection),
	unregister: make(chan *connection),
	connections: make(map[*connection]bool),
	pConnections: make(map[string]*connection),
	pConnectionsSet: make(map[*connection]string),
	pRegister: make(chan PlayerPageRegistration),
	updatedUids: make(chan string),

	
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
		case uid := <-h.updatedUids:
			if val,ok := h.pConnections[uid]; ok {
				updatePlayerPage(val, uid)
			}
		case s := <-h.pRegister:
			fmt.Println("New client connected to a player page.")
			h.pConnections[s.Uid] = s.Conn
			h.pConnectionsSet[s.Conn] = s.Uid
			updatePlayerPage(s.Conn, s.Uid)
		case c := <-h.pUnregister:
			fmt.Println("Client disconnected from player page")
			delete(h.pConnections, h.pConnectionsSet[c])
			delete(h.pConnectionsSet, c)
		}
	}
}


