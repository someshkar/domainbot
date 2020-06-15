package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/someshkar/domainbot/handlers"
)

func main() {
	// Load env variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	discordToken := os.Getenv("DISCORD_TOKEN")

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Add handlers
	dg.AddHandler(handlers.MainHandler)

	// Open a websocket connection to Discord and start listening
	err = dg.Open()
	if err != nil {
		log.Fatalln("Error opening connection", err)
		return
	}

	// Wait here until Ctrl-C or some other term signal is received
	fmt.Println("Domainbot is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close the Discord session
	dg.Close()
}
