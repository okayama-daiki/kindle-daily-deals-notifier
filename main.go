package main

import (
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/crawler"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/notifier"
)

var (
	targetId = os.Getenv("LINE_TARGET_ID")
)

func formatMessage(productName string, productUrl url.URL) string {
	return productName + "\n" + productUrl.String()
}

func lambdaHandler(req events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	productList, err := crawler.Crawl()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	messages := []messaging_api.MessageInterface{}
	for _, product := range productList {
		message := messaging_api.TextMessage{
			Text: formatMessage(product.Name, product.URL),
		}
		messages = append(messages, message)
	}

	notifier.Notify(targetId, messages)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
