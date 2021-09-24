package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"time"
)

/*
	AV - Audio/Video
*/

const defaultLimit = 3
var (
	responses = [3]string{
		"Пиши буковами, %s",
		"Я запрещаю вам срать, %s",
		"Президент заборонив срати, %s",
	}
	cache = make(map[int64]map[int]Cache)
)
type Cache struct {
	time time.Time
	sendVoice int64
}


func updateCache(chatID int64, update tgbotapi.Update) {
	if _, ok := cache[chatID]; !ok {
		cache[chatID] = map[int]Cache{
			update.Message.From.ID: {
				time.Now(),
				1,
			},
		}
	} else {
		if _, ok := cache[chatID][update.Message.From.ID]; !ok {
			cache[chatID][update.Message.From.ID] = Cache{
				time.Now(),
				1,
			}
		} else {
			cached := cache[chatID][update.Message.From.ID]
			if time.Now().Sub(cached.time).Hours() >= 24 {
				cache[chatID][update.Message.From.ID] = Cache{
					time.Now(),
					1,
				}
			} else if cached.sendVoice <= defaultLimit {
				cache[chatID][update.Message.From.ID] = Cache{
					time.Now(),
					cached.sendVoice + 1,
				}
			}
		}
	}
}


func HandleAV(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID

	updateCache(chatID, update)

	if cache[chatID][update.Message.From.ID].sendVoice <= defaultLimit {
		return
	}

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
