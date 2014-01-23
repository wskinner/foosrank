import (
	"fmt"
	"html/template"
)

type Opponent struct {
	WinsAgainst int
	LossesAgainst int
	Player Player
}

type PlayerData struct {
	Player RankedPlayer
	Opponents [] // 
}

func playersTemplate(uid string) (templ html.Template, data PlayerData) {
	templ := template.Must(template.ParseFiles(filepath.Join(defaultAssetPath(), "players.html")))
	data := // get games from db

}
