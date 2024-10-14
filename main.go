package main

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/okayama-daiki/kindle-daily-deals-notifier/libs/lambda"
)

type Config struct {
	ChannelAccessToken string `env:"LINE_CHANNEL_ACCESS_TOKEN"`
	TargetId           string `env:"LINE_TARGET_ID"`
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}

	lambda, err := lambda.New(cfg.ChannelAccessToken, cfg.TargetId)
	if err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}

	lambda.Start()
}
