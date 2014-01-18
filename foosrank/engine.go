package foosrank

import (
	"flag"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
	"encoding/json"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	assets = flag.String("assets", defaultAssetPath(), "path to assets")
	homeTempl *template.Template
)

func defaultAssetPath() string {
	p, err := build.Default.Import("github.com/wskinner/foosrank/foosrank", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

func Run() {
	flag.Parse()
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, "home.html")))
	go h.run()
	
	// Just feed some jsons in there every 5 seconds
	go func() {
		p1 := Player{"Will", "Skinner"}
		r1 := EloRank{1500}
		rp1 := RankedPlayer{p1, r1}

		p2 := Player{"Michael", "Schiff"}
		r2 := EloRank{1500}
		rp2 := RankedPlayer{p2, r2}

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
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
