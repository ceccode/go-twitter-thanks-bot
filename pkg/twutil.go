package pkg

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

// Sends a direct message
func SendDirectMessage(client *twitter.Client, screenName string, messageText string) {
	dmParams := &twitter.DirectMessageNewParams{
		ScreenName: screenName,
		Text:       messageText,
	}
	directMessage, httpResponse, err := client.DirectMessages.New(dmParams)
	if err != nil {
		fmt.Println(directMessage, httpResponse, err)
	}
}
