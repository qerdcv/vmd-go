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
	statistics := GetStatisticForChat(c.Update.Message.Chat.ID)
	if len(statistics) == 0 {
		msg := tgbotapi.NewMessage(c.Update.Message.Chat.ID, "There is no statistic yet")
		_, err := c.Bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
		return
	}

	resp += fmt.Sprintf("https://stat.qerdcv.com?chat_id=%d\n", c.Update.Message.Chat.ID)
	for idx, statistic := range statistics {
		resp += fmt.Sprintf(
			"%d. %s настрал %d ",
			idx + 1,
			statistic.Username,
			statistic.DeletedCount)
		resp += "раз(а)\n"
	}

	msg := tgbotapi.NewMessage(c.Update.Message.Chat.ID, resp)
	_, err := c.Bot.Send(msg)

	if err != nil {
		log.Println(err)
	}
}
