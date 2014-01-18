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

type Leaderboard struct {
	Players [2]string
}

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
		players := [2]string{"Will", "Michael"}
		msg, _ := json.Marshal(Leaderboard {Players: players})
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
