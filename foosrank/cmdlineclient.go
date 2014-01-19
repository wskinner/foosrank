package foosrank

import (
    "os"
    "fmt"
    "bufio"
    "github.com/wskinner/anaconda"
)

func ReadTweetFromStdIn(tweetChan chan anaconda.Tweet) {
    bio := bufio.NewReader(os.Stdin)
    for {
        line, _, err := bio.ReadLine()
        if err == nil {
            text := string(line)
            fmt.Println(text)
            tweet := anaconda.Tweet{}
            tweet.Text = text
            tweetChan <- tweet
        }
    }
}
