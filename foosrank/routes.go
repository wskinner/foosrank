package foosrank

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	//"strings"
	"go/build"
	"path/filepath"
	"html/template"
)


func defaultAssetPath() string {
	p, err := build.Default.Import("github.com/wskinner/foosrank/foosrank/assets", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	homeTempl := template.Must(template.ParseFiles(filepath.Join(defaultAssetPath(), "home.html")))
	fmt.Println("Host: " + r.Host)
	homeTempl.Execute(w, r.Host)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}

	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c
	defer func() { h.unregister <-c }()
	go c.writer()
	c.reader()
}
/*j
func playersHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	split := strings.Split(url.Path, "/")
	uid := split[len(split)-1]
	fmt.Println("uid: ", uid)
	//
	playerTempl, data := playersTemplate(uid)

	playerTempl.Execute(w, data)
}
*/

