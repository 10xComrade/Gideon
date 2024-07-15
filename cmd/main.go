package main

import (
	"Gideon/actions"
	"Gideon/bot"
	"Gideon/config"
	"Gideon/handler"
	"log"
)

// main function
func main() {
	config.ReadConfig()

	bot, err := bot.Init()

	if err != nil {
		log.Fatal("Could not initiate the bot !")
		return
	}

	log.Printf("Authorized on account : %s", bot.Self.UserName)

	newsReader := actions.NewNewsReader(config.GlobalConfig.NewsAPIToken)

	log.Println("Beep boop Beep zhzhzh ...")
	log.Println("Receiving Updates ...")

	handler.HandleUpdates(bot, newsReader)
}
