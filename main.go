package main

import (
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func pushToUser(message *TiDBMessage) error {

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		return err
	}

	m, err := NewFlex(message.PhotoURL, message.Message)
	if err != nil {
		return err
	}
	_, err = bot.PushMessage("U316e79437f87b71ce6d0965e8ab47453", m).Do()
	if err != nil {
		return err
	}

	return nil
}

func HandleRequest(event LambdaFunctionURLRequest) (string, error) {
	slog.Info("Start")

	const ID = 60001
	m, err := FindMessage(ID)
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
