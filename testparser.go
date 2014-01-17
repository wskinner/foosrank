package main

import (
	"github.com/wskinner/foosrank/foosrank"
)

func main() {
	foosrank.GetTweetEntities("will skinner 6 michael schiff 8")
	foosrank.GetTweetEntities("will 6 michael schiff 8")
	foosrank.GetTweetEntities("will skinner 6 michael 8")
	foosrank.GetTweetEntities("will 6 michael 8")
}
