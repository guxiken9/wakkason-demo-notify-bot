package main

import (
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() (string, error) {
	slog.Info("Start")

	// 投稿するメッセージを取得

	// メッセージ投稿

	slog.Info("End")
	return "### success ###", nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	lambda.Start(HandleRequest)
}
