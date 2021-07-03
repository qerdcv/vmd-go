package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
)

/*
	AV - Audio/Video
*/

var responses = [3]string{
	"Пиши буковами, %s",
	"Я запрещаю вам срать, %s",
	"Президент заборонив срати, %s",
}

func HandleAV(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID

	_, err := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID: chatID,
		MessageID: messageID,
	})
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Failed to delete message")
		log.Println(bot.Send(msg))
		return
	}
	fromUser := fmt.Sprintf("@%s", update.Message.From.UserName)
	response := fmt.Sprintf(responses[rand.Intn(len(responses))], fromUser)
	msg := tgbotapi.NewMessage(chatID, response)
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("Error while sending message")
	}
	err = createStatistic(
		chatID,
		update.Message.From.ID,
		update.Message.From.UserName)
	if err != nil {
		// Statistic for that chat and user id is already exists
		updateStatistic(
			chatID,
			update.Message.From.ID)
	}
}
