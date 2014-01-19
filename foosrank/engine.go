package foosrank

import (
	"flag"
	"log"
	"net/http"
	"encoding/json"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
)

func RunServer(leaderboardChan chan []RankedPlayer) {
	flag.Parse()
	go h.run()
	
    go func() {
	    http.HandleFunc("/", homeHandler)
	    http.HandleFunc("/ws", wsHandler)
	    http.HandleFunc("/players/", playersHandler)
	    if err := http.ListenAndServe(*addr, nil); err != nil {
		    log.Fatal("ListenAndServe:", err)
	    }
        for leaderboard := range leaderboardChan {
		    msg, _ := json.Marshal(leaderboard)
            h.broadcast <- msg
	    }
    }()
}
