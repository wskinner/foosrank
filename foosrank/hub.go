package foosrank

type hub struct {
	// Registered connections
	connections map[*connection]bool

	// Inbound messages from the connections
	broadcast chan []byte

	// Register requests from the connections
	register chan *connection

	// Unregister requests from the connections
	unregister chan *connection
}

var h = hub {
	broadcast: make(chan [] byte),
	register: make(chan *connection),
	unregister: make(chan *connection),
	connections: make(map[*connection]bool),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			delete(h.connections, c)
		case m := <-h.broadcast:
			// For each connection, send it the message. If the channel is not 
			// full, close the connection.
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


