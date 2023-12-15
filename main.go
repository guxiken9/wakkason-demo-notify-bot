package main

import (
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func pushToUser() error {
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		return err
	}
	res, err := bot.PushMessage(
		&messaging_api.PushMessageRequest{
			To: "annkara",
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: "Test from Lambda",
				},
			},
		},
		"", // x-line-retry-key
	)
	if err != nil {
		return err
	}

	slog.Info("Push Message Response", res)

	return nil
}

func HandleRequest() (string, error) {
	slog.Info("Start")

	// 投稿するメッセージを取得

	// メッセージ投稿
	if err := pushToUser(); err != nil {
		slog.Error("ユーザへの通知にエラーとなりました。", err)
		return "", err
	}

	slog.Info("End")
	return "### success ###", nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	lambda.Start(HandleRequest)
}
