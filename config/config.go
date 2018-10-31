package config

import "github.com/BurntSushi/toml"

type ConfigData struct {
	Slack Slack
	AWS   AWS
}

type Slack struct {
	SendAPIURL    string
	ChannelName   string
	RegsiterToken string
}

type AWS struct {
	AccessKey string
	SecretKey string
	Region    string
	DataBase  string
}

/**
 * 設定データを取得する
 */
func GetConfigData() ConfigData {
	var configData ConfigData
	_, err := toml.DecodeFile("config.toml", &configData)
	if err != nil {
		panic("tomlファイルを読み込めません")
	}

	return configData
}
