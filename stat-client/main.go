package stat_client

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func setupHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GEGE")
	}))
	return mux
}

func Run() {
	handler := setupHandler()
	server := &http.Server{
		Addr: ":3002",
		Handler: handler,
		ReadTimeout: 60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
