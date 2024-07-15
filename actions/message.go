package actions

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendRawMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}

// TODO : make this function more flexible
func SendMessageWithKeyboard(bot *tgbotapi.BotAPI, chatID int64, text string) {
	var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Previous", "news_btn_left"),
			tgbotapi.NewInlineKeyboardButtonData("Read", "news_btn_read"),
			tgbotapi.NewInlineKeyboardButtonData("Next", "news_btn_right"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "")
	msg.ReplyMarkup = inlineKeyboard
	msg.Text = text

	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

// TODO : make a addline function with string builder seperately
func WrapArticleAsMessage(article *Article, currentPage int) string {
	var builder strings.Builder

	addLine := func(label, value string) {
		if value != "" {
			builder.WriteString("\n" + label + value + "\n")
		}
	}

	addLine("🏷️ 𝐓𝐢𝐭𝐥𝐞 : ", article.Title)
	addLine("📅 𝐏𝐮𝐛𝐥𝐢𝐬𝐡𝐞𝐝 𝐚𝐭 : ", article.PublishedAt)
	addLine("📝 𝐃𝐞𝐬𝐜𝐫𝐢𝐩𝐭𝐢𝐨𝐧 : ", article.Description)
	addLine("🔗 𝐔𝐑𝐋 : ", article.URL)
	addLine("🖼️ 𝐔𝐑𝐋𝐓𝐨𝐈𝐦𝐚𝐠𝐞 : ", article.URLToImage)
	addLine("📖 𝐑𝐞𝐬𝐮𝐥𝐭 : ", fmt.Sprintf("%d 𝐟𝐫𝐨𝐦 %d",
		currentPage, len(GlobalNewsResponse.Articles)-1))

	return builder.String()
}
