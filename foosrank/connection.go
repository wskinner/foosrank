package foosrank

import (
	"github.com/gorilla/websocket"
	"fmt"
	"encoding/json"
)

type connection struct {
	// The websocket connection
	ws *websocket.Conn

	// Buffered channel of outbound messages
	send chan []byte
}

type Pong struct {
	Pong string
}

// Handle pongs
func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		var msg Pong
		_ = json.Unmarshal(message, &msg)
		fmt.Println("Received message:", msg)

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

