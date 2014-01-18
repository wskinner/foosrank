package foosrank

import (
	"github.com/gorilla/websocket"
)

type connection struct {
	// The websocket connection
	ws *websocket.Conn

	// Buffered channel of outbound messages
	send chan []byte
}

// For now this is a nop
func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		var _ = message

		//h.broadcast <- message
	}
	// Getting here means an error occurred
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

