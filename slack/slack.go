package slack

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"strings"

	config "../config"
)

/**
 * Slackからのパラメータ格納構造体
 */
type SlackParams struct {
	Token    string
	UserID   string
	UserName string
	Text     string
}

/**
 * Slackの特定チャンネルにメッセージを投稿する
 */
func SendMessage(configData config.ConfigData, sendText string) {
	// JSONとしてパラメータを設定
	jsonStr := `{"text":"` + sendText + `","channel":"` + configData.Slack.ChannelName + `","link_names":"1"}`

	// 通知を実行する
	request, newRequestError := http.NewRequest(
		"POST",
		configData.Slack.SendAPIURL,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if newRequestError != nil {
		panic("newRequestErrorのリクエスト作成に失敗しました。")
	}
	defer request.Body.Close()

	request.Header.Add("Content-pe", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, doSendError := client.Do(request)
	if doSendError != nil {
		panic("SendMessageのリクエスト実行に失敗しました。")
	}

	defer response.Body.Close()
}

/**
 * Slackから送られたパラメータをパースする
 * @see https://qiita.com/holy_road_ss/items/51f988174be8d39e9c5f#golanglambdaslack%E9%83%A8%E5%88%86
 */
func ParseSlackParams(rawParams interface{}) (result *SlackParams, err error) {
	result = &SlackParams{}

	tmp := rawParams.(map[string]interface{})
	if _, ok := tmp["body"]; !ok {
		err = errors.New("params body does not exists")
		return
	}
	rawQueryString := tmp["body"].(string)
	parsed, err := url.QueryUnescape(rawQueryString)
	if err != nil {
		err = errors.New("params body unescape failed. body: " + rawQueryString)
		return
	}
	params, err := url.ParseQuery(parsed)
	if err != nil {
		err = errors.New("params body parse failed. body: " + rawQueryString)
		return
	}

	result.Token = params["token"][0]
	result.UserID = params["user_id"][0]
	result.UserName = params["user_name"][0]
	result.Text = params["text"][0]

	// Slash CommandでURL形式を送ると <URL>という形式になるので、先頭と末尾をtrimする
	result.Text = strings.TrimLeft(result.Text, "<")
	result.Text = strings.TrimRight(result.Text, ">")

	return
}
