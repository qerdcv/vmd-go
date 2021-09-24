package main

import (
	"math/rand"
	"time"
	"vmd-go/bot"
	client "vmd-go/stat-client"
)

func main() {
	rand.Seed(time.Now().Unix())
	go bot.RunBot()
	client.Run()
}
