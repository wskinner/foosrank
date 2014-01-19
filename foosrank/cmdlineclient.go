package foosrank

import (
    "fmt"
    "github.com/wskinner/anaconda"
)

func ReadTweetFromStdIn(tweetChan chan anaconda.Tweet) {
    for {
        var text string
        _, err := fmt.Scanf("%s", &text)
        if err == nil {
            tweet := anaconda.Tweet{}
            tweet.Text = text
            tweetChan <- tweet
        }
    }
}
