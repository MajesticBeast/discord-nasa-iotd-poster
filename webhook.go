package main

import (
	"log"

	"github.com/gtuk/discordwebhook"
)

func publishToHook(t string, c string) {

	message := discordwebhook.Message{
		Username: &t,
		Content:  &c,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}
