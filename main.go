package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
  "github.com/ceccode/go-twitter-thanks-bot/pkg"
)

func configure() {

	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
	}
	
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	
	demux := twitter.NewSwitchDemux()	

	demux.Event = func(event *twitter.Event){
		fmt.Printf("%#v\n", event)
		fmt.Println(event.Event)
		if (event.Event == "follow") {
			pkg.SendDirectMessage(client, event.Source.ScreenName , "@" + event.Source.ScreenName + " Thank you for the follow up.")
		}
	}
	
	fmt.Println("Starting Stream...")

	stream, err := client.Streams.User(&twitter.StreamUserParams{})

	if err != nil {
		log.Fatal(err)
	}
	
	go demux.HandleChan(stream.Messages)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
	
}

func main() {
	fmt.Println("Go-Twitter Thanks Bot v0.01")
	configure()
}
