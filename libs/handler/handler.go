package handler

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/crawler"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/notifier"
)

type Handler struct {
	bot      *messaging_api.MessagingApiAPI
	targetId string
}

func New(bot *messaging_api.MessagingApiAPI, targetId string) *Handler {
	return &Handler{
		bot,
		targetId,
	}
}

func (h *Handler) HandleRequest(req events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
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

	notifier := notifier.New(h.bot)
	if err := notifier.Notify(h.targetId, messages); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}
