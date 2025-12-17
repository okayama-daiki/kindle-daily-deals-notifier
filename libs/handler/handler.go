package handler

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/crawler"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/notifier"
)

func Handler(bot *messaging_api.MessagingApiAPI, targetId string) func(events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	return func(req events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
		products, err := crawler.Crawl()
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
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
		for i := 0; i < len(messages); i += 5 {
			// NOTE: LINE Messaging API allows maximum 5 messages per request
			if err := notifier.Notify(targetId, messages[i:i+5]); err != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
				}, err
			}
		}

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
		}, nil
	}

}
