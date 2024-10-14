package notifier

import (
	"log"

	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type Notifier struct {
	bot *messaging_api.MessagingApiAPI
}

func New(bot *messaging_api.MessagingApiAPI) *Notifier {
	return &Notifier{
		bot: bot,
	}
}

func (n *Notifier) Notify(targetId string, messages []messaging_api.MessageInterface) error {
	retryKey := uuid.NewString()

	response, err := n.bot.PushMessage(
		&messaging_api.PushMessageRequest{
			To:       targetId,
			Messages: messages,
		},
		retryKey,
	)

	if err != nil {
		return err
	}

	log.Println("Pushed message to", targetId)
	log.Println("Pushed response:", response)

	return nil
}
