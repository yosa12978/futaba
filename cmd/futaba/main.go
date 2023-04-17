package main

import (
	"github.com/joho/godotenv"
	bot "github.com/yosa12978/futaba/internal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	bot.Run()
}
