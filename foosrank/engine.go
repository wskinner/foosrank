package foosrank

import (
	"flag"
	"log"
	"net/http"
	"encoding/json"
)

var (
	addr = flag.String("addr", ":80", "http service address")
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
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/players/", playersHandler)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
}
}
