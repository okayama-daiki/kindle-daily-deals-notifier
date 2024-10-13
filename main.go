package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env/v11"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/crawler"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/notifier"
)

type config struct {
	ChannelAccessToken string `env:"LINE_CHANNEL_ACCESS_TOKEN"`
	TargetId           string `env:"LINE_TARGET_ID"`
}

var (
	cfg config
	bot *messaging_api.MessagingApiAPI
)

func init() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	var err error
	bot, err = messaging_api.NewMessagingApiAPI(
		cfg.ChannelAccessToken,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func lambdaHandler(req events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	products, err := crawler.Crawl()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	var messages []messaging_api.MessageInterface
	for _, product := range products {
		message := messaging_api.TextMessage{
			Text: product.String(),
		}
		messages = append(messages, message)
	}

	notifier := notifier.New(bot)
	notifier.Notify(cfg.TargetId, messages)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
