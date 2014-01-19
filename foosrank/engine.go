package foosrank

import (
	"flag"
	"log"
	"net/http"
	"time"
	"encoding/json"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
)

func Run() {
	flag.Parse()
	go h.run()
	
	// Just feed some jsons in there every 5 seconds
	go func() {
		p1 := Player{"Will", "Skinner"}
		r1 := EloRank{1500}
		rp1 := RankedPlayer{p1, r1}//, 1234}

		p2 := Player{"Michael", "Schiff"}
		r2 := EloRank{1500}
		rp2 := RankedPlayer{p2, r2}//, 5678}

		leaderboard := Leaderboard{[]RankedPlayer{rp1, rp2}}
		msg, _ := json.Marshal(leaderboard)
		dur, _ := time.ParseDuration("5s")
		for {
			h.broadcast <- msg
			time.Sleep(dur)
		}
	}()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/players/", playersHandler)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
