package foosrank

import (
	"flag"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"time"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
)

func RunServer(leaderboardChan chan []RankedPlayer) {
	flag.Parse()
	go h.run()
	
	go func() {
		for leaderboard := range leaderboardChan {
			msg, _ := json.Marshal(leaderboard)
			h.currentLeaderboard = leaderboard
			h.broadcast <- msg
		}
	}()

	go func() {
		for {
			fmt.Printf("There are %d clients connected\n", len(h.connections))
			t, _ := time.ParseDuration("30s")
			time.Sleep(t)
		}
	}()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/pws/", playersWsHandler)
	http.HandleFunc("/players/", playersHandler)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
