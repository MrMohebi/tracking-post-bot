package myTel

import (
	"bytes"
	"encoding/json"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/configs"
	"net/http"
)

func sendReq(urlPath string, data interface{}) bool {
	baseURL := configs.EnvTelBaseUrl()
	dataJSON, err := json.Marshal(data)
	common.IsErr(err)
	_, err = http.Post(baseURL+urlPath, "application/json", bytes.NewBuffer(dataJSON))
	common.IsErr(err)
	return true
}

func SendText(text string, chatId int64) bool {
	type req struct {
		Token  string `json:"token"`
		Text   string `json:"text"`
		ChatId int64  `json:"chatId,string"`
	}

	sendReq("sendText", req{
		Token:  configs.EnvBotToken(),
		Text:   text,
		ChatId: chatId,
	})

	return true
}
