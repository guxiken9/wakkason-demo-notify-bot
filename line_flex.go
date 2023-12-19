package main

import (
	"encoding/json"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func UnmarshalUser(data []byte) (Message, error) {
	var r Message
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Message) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Message struct {
	Type string `json:"type"`
	H    Hero   `json:"hero"`
	B    Body   `json:"body"`
}

type Hero struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Size        string `json:"size"`
	AspectRatio string `json:"aspectRatio"`
	AspectMode  string `json:"aspectMode"`
}

type Body struct {
	BodyType string  `json:"type"`
	Layout   string  `json:"layout"`
	Contents Content `json:"contents"`
}

type Content struct {
	ContentType string `json:"type"`
	Text        string `json:"text"`
	Weight      string `json:"weight"`
	Size        string `json:"size"`
}

const jsonData1 = `{
  "type": "bubble",
  "hero": {
    "type": "image",
    "url": `
const jsonData2 = `,
    "size": "full",
    "aspectRatio": "20:13",
    "aspectMode": "cover"
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": `
const jsonData3 = `,
        "weight": "bold",
        "size": "xl"
      }
    ]
  }
}`

func NewFlex(url, text string) (*linebot.FlexMessage, error) {

	jsonStr := jsonData1 + url + jsonData2 + text + jsonData3
	container, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonStr))
	if err != nil {
		// 正しくUnmarshalできないinvalidなJSONであればerrが返る
		return nil, err
	}
	message := linebot.NewFlexMessage("alt text", container)

	return message, nil
}
