package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/crawler"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/notifier"
)

var (
	channelAccessToken = os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	targetId           = os.Getenv("LINE_TARGET_ID")
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
	notifier.Notify(targetId, messages)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
