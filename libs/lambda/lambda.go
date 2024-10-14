package lambda

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/handler"
)

type Lambda struct {
	bot      *messaging_api.MessagingApiAPI
	targetId string
}

func New(channelAccessToken string, targetId string) (*Lambda, error) {
	bot, err := messaging_api.NewMessagingApiAPI(channelAccessToken)
	if err != nil {
		return nil, err
	}

	return &Lambda{
		bot,
		targetId,
	}, nil
}

func (l *Lambda) Start() {
	lambda.Start(handler.Handler(l.bot, l.targetId))
}
