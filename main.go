package main
import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"eef.gocord/bot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	botToken := os.Getenv("DISCORD_TOKEN")
	bot.Token = botToken;
	bot.Start();
}