package stat_client

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"vmd-go/bot"
)

func statisticHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index.gohtml").Funcs(funcMap).ParseFiles("./ui/index.gohtml")
	if err != nil {
		log.Println(err)
	}
	if chatID := r.URL.Query().Get("chat_id"); chatID != "" {
		// Statistic for chatID
		chatID, _ := strconv.ParseInt(chatID, 10, 64)
		err = tmpl.Execute(w,  &Page{
			Statistics: bot.GetStatisticForChat(
				chatID),
		})
	} else {
		// Whole statistic
		// TODO: add pagination
		err = tmpl.Execute(w,  &Page{
			Statistics: bot.GetStatistic(),
		})
	}
	if err != nil {
		log.Println(err)
	}
}

func SetupHandlers() http.Handler {
	mux := http.NewServeMux()

	// Static
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))
	// Statistic handler
	mux.Handle("/", http.HandlerFunc(statisticHandler))
	return mux
}