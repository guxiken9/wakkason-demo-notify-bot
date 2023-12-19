package main

import (
	"encoding/json"
)

type NotifyRequest struct {
	MessageID string `json:"messageId"`
}

// Lambda Function URL のリクエスト
type LambdaFunctionURLRequest struct {
	Body string `json:"body"`
}

func UnmarshalLambdaRequestBody(data []byte) (NotifyRequest, error) {
	var r NotifyRequest
	err := json.Unmarshal(data, &r)
	return r, err
}
