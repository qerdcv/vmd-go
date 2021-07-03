package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Command struct {
	Bot *tgbotapi.BotAPI
	Update tgbotapi.Update
}

func (c *Command) Stat() {
	var resp string
	statistics := getStatistic(c.Update.Message.Chat.ID)
	if len(statistics) == 0 {
		msg := tgbotapi.NewMessage(c.Update.Message.Chat.ID, "There is no statistic yet")
		_, err := c.Bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
		return
	}

	for idx, statistic := range statistics {
		resp += fmt.Sprintf(
			"%d. %s настрал %d ",
			idx + 1,
			statistic.Username,
			statistic.DeletedCount)
		if statistic.DeletedCount < 5  && statistic.DeletedCount > 1 {
			resp += "раза"
		} else {
			resp += "раз"
		}
		resp += "\n"
	}

	msg := tgbotapi.NewMessage(c.Update.Message.Chat.ID, resp)
	_, err := c.Bot.Send(msg)

	if err != nil {
		log.Println(err)
	}
}
