package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	slackBotToken := goDotEnvVariable("SLACK_BOT_TOKEN")
	channelID := goDotEnvVariable("CHANNEL_ID")

}

func goDotEnvVariable(key string) string {
	//load .env
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
