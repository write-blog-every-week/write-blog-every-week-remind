package config

import "github.com/BurntSushi/toml"

// ConfigData Base
type ConfigData struct {
	Slack Slack
	AWS   AWS
}

// Slack Slack用設定データ格納構造体
type Slack struct {
	SendAPIURL  string
	ChannelName string
}

// AWS AWS用設定データ格納構造体
type AWS struct {
	AccessKey string
	SecretKey string
	Region    string
	DataBase  string
}

// GetConfigData 設定データを取得する
func GetConfigData() ConfigData {
	var configData ConfigData
	_, err := toml.DecodeFile("config.toml", &configData)
	if err != nil {
		panic("tomlファイルを読み込めません")
	}

	return configData
}
