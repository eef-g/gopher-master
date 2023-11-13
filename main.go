package main

import (
	"fmt"
	"os"

	"eef.gocord/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	botToken := os.Getenv("DISCORD_TOKEN")
	fmt.Println(botToken)
	bot.Token = botToken
	bot.Start()
}
