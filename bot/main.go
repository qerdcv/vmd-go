package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

const (
	Stat = "stat"
)

func InitBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	return bot, updates
}

func RunBot() {
	bot, updates := InitBot()
	for update := range updates {
		c := &Command{
			Bot: bot,
			Update: update,
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case Stat:
				c.Stat()
				break
			}
		} else if update.Message.Voice != nil || update.Message.VideoNote != nil {
			HandleAV(bot, update)
		}
	}
	log.Println("init bot")
}
