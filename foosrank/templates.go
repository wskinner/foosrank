package foosrank

import (
//	"html/template"
//	"path/filepath"
)

type Opponent struct {
	WinsAgainst int
	LossesAgainst int
	Player Player
}

type PlayerData struct {
	Player RankedPlayer
	Opponents []Opponent // 
}
/*
func playersTemplate(uid string) (templ template.Template, data PlayerData) {
	templ = template.Must(template.ParseFiles(filepath.Join(defaultAssetPath(), "players.html")))
	//data = // get games from db
	return nil, nil
}
*/
