package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	slackBotToken := goDotEnvVariable("SLACK_BOT_TOKEN")
	channelID := goDotEnvVariable("CHANNEL_ID")

	api := slack.New(slackBotToken)

	fileInfo, err := os.Stat("Testing_Jazz.pdf")

	if err != nil {
		fmt.Printf("Error parsing file: %s\n", err)
		return
	}

	//Flattened this out of an array as prior technique didn't factor in new method and its params.
	params := slack.UploadFileV2Parameters{
		Channel:  channelID,
		File:     fileInfo.Name(),
		Filename: fileInfo.Name(),
		FileSize: int(fileInfo.Size()),
	}

	file, err := api.UploadFileV2(params)

	if err != nil {
		fmt.Printf("Error uploading file: %s\n", err)
		return
	}

	fmt.Printf("File uploaded:%s\n", file)

}

func goDotEnvVariable(key string) string {
	//load .env
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
