package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func main() {
	discord, err := discordgo.New("Bot " + "")
	if err != nil {
		log.Fatal("ggwp")
	}

}
