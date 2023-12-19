package main

import (
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func pushToUser(message *TiDBMessage) error {
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		return err
	}

	_, err = bot.PushMessage(
		&messaging_api.PushMessageRequest{
			To: "U319905930de67669d4d53848cd3325a1",
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: message.Title,
				},
			},
		},
		"", // x-line-retry-key
	)
	if err != nil {
		return err
	}

	return nil
}

func HandleRequest(event LambdaFunctionURLRequest) (string, error) {
	slog.Info("Start")

	// 投稿するメッセージを取得
	r, err := UnmarshalLambdaRequestBody([]byte(event.Body))
	if err != nil {
		return "", err
	}

	m, err := FindMessage(r.MessageID)
	if err != nil {
		slog.Error("ユーザへの通知にエラーとなりました。", err)
		return "", err
	}

	// メッセージ投稿
	if err := pushToUser(m); err != nil {
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
