package handler

import (
	"Gideon/actions"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUpdates(bot *tgbotapi.BotAPI, newsReader *actions.NewsReader) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 5

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal("Error setting up updates:", err)
		return
	}

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			HandleCommands(bot, newsReader, &update)

		} else if update.CallbackQuery != nil {
			HanldeNewsMessageCallBackQueries(bot, newsReader, &update)
		}
	}
}

func HandleCommands(bot *tgbotapi.BotAPI,
	newsReader *actions.NewsReader,
	update *tgbotapi.Update,
) {
	switch update.Message.Command() {

	case "start":
		StartCommand(bot, update)

	case "help":
		HelpCommand(bot, update)

	case "news":
		NewsCommand(bot, newsReader, update)
	}
}

func HanldeNewsMessageCallBackQueries(bot *tgbotapi.BotAPI,
	newsReader *actions.NewsReader,
	update *tgbotapi.Update,
) {
	if actions.GlobalNewsResponse == nil {
		return
	}

	currentPage := actions.CurrentPage

	switch update.CallbackQuery.Data {
	case "news_btn_left":
		*currentPage--

	case "news_btn_read":
		url := actions.GlobalNewsResponse.Articles[*currentPage].URL
		content := actions.Scrape('p', url, 4091)
		actions.SendRawMessage(bot, update.CallbackQuery.Message.Chat.ID, content)
		return

	case "news_btn_right":
		*currentPage++
	}

	actions.LimitCurrentPage(currentPage)
	articles := actions.GlobalNewsResponse.Articles

	msg := actions.WrapArticleAsMessage(&articles[*currentPage], *currentPage)
	actions.SendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID, msg)
}

func NewsCommand(bot *tgbotapi.BotAPI,
	newsReader *actions.NewsReader,
	update *tgbotapi.Update,
) {
	args := strings.Split(update.Message.CommandArguments(), " ")
	if len(args) < 3 {
		text := "Gideon requires exactly 3 arguments."
		actions.SendRawMessage(bot, update.Message.Chat.ID, text)
		return
	}

	currentPage, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		text := "Please enter correct format of page number !"
		actions.SendRawMessage(bot, update.Message.Chat.ID, text)
		return
	}

	subject := strings.Join(args[:len(args)-2], " ")
	sortBy := args[len(args)-2]

	news, err := newsReader.FetchNews(subject, sortBy)

	if err != nil || len(news.Articles) == 0 {
		text := "No articles found !"
		actions.SendRawMessage(bot, update.Message.Chat.ID, text)
		return
	}

	articles := news.Articles
	actions.LimitCurrentPage(&currentPage)

	msg := actions.WrapArticleAsMessage(&articles[currentPage], currentPage)
	actions.SendMessageWithKeyboard(bot, update.Message.Chat.ID, msg)
}

func StartCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	startMessage := `
Greetings, i'm Gideon !
a Telegram bot which is
inspired by Gideon the AI in the Flash Series.

My duty is to gather latest news
,enlighten people & show them what's
happening around the world !

If you had any questions please
contact my Developer : @CipherChunk

Feel free to share issues, contribute
to the project : 
github.com/10xComrade

These commands are known to me : 
/start , /help , /news 
	`

	actions.SendRawMessage(bot, update.Message.Chat.ID, startMessage)
}

func HelpCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	helpMessage := `
/start - start the bot 
/help - show help
/news - show news  

Usage : 
/news <SUBJECT> <SORT-BY> <PAGE-NUM>

Examples : 
1) /news tesla publishedAt 0 
2) /news bitcoin relevancy 2
	`

	actions.SendRawMessage(bot, update.Message.Chat.ID, helpMessage)
}
