package bot

import (
	"Gideon/config"
	"Gideon/handler"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Init() (*tgbotapi.BotAPI, error) {
	botToken := config.GlobalConfig.BotToken

	var bot *tgbotapi.BotAPI
	var err error

	if config.GlobalConfig.Proxy.Enabled {
		proxyURL := config.GlobalConfig.Proxy.URL
		proxyHandler := handler.NewProxyHandler()

		httpClient, errP := proxyHandler.NewProxyClient(proxyURL)
		if errP != nil {
			log.Println("Error creating proxy client:", err)
			return nil, err
		}

		bot, err = tgbotapi.NewBotAPIWithClient(botToken, httpClient)

	} else {
		bot, err = tgbotapi.NewBotAPI(botToken)
	}

	if err != nil {
		return bot, err
	}

	return bot, nil
}
