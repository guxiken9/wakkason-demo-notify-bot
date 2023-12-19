package main

import (
	"encoding/json"
)

type LambdaFunctionURLRequest struct {
	Body string `json:"body"`
}

func UnmarshalWelcome(data []byte) (LambdaFunctionURLRequest, error) {
	var r LambdaFunctionURLRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

type NotifyRequest struct {
	MessageID int `json:"messageId"`
}

func UnmarshalLambdaRequestBody(data []byte) (NotifyRequest, error) {
	var r NotifyRequest
	err := json.Unmarshal(data, &r)
	return r, err
}
