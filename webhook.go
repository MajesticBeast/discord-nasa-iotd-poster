package main

import (
	"fmt"
	"log"

	"github.com/gtuk/discordwebhook"
)

func publishToHook(title string, desc string, content string) {

	content = fmt.Sprintf("%s\n\n%s", desc, content)

	message := discordwebhook.Message{
		Username: &title,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}
