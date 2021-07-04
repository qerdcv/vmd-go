package stat_client

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"vmd-go/bot"
)


type Page struct {
	Statistics []bot.Statistic
}

var funcMap = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
}


func Run() {
	handler := SetupHandlers()
	server := &http.Server{
		Addr:         ":3001",
		Handler:      handler,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
