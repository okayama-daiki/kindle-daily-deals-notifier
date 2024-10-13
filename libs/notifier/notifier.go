package notifier

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

var (
	channelAccessToken = os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	bot                *messaging_api.MessagingApiAPI
)

func init() {
	var err error
	bot, err = messaging_api.NewMessagingApiAPI(
		channelAccessToken,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func Notify(targetId string, messages []messaging_api.MessageInterface) {
	retryKey := uuid.NewString()

	response, err := bot.PushMessage(
		&messaging_api.PushMessageRequest{
			To:       targetId,
			Messages: messages,
		},
		retryKey,
	)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Pushed message to", targetId)
		log.Println("Pushed response:", response)
	}
}
