package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env/v11"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/handler"
)

type config struct {
	ChannelAccessToken string `env:"LINE_CHANNEL_ACCESS_TOKEN"`
	TargetId           string `env:"LINE_TARGET_ID"`
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	var cfg config

	if err := env.Parse(&cfg); err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}

	bot, err := messaging_api.NewMessagingApiAPI(cfg.ChannelAccessToken)
	if err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}

	handler := handler.New(bot, cfg.TargetId)
	lambda.Start(handler.HandleRequest)
}
