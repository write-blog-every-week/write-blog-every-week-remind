package config

import "github.com/BurntSushi/toml"

// ConfigData Base
type ConfigData struct {
	Slack Slack
	AWS   AWS
	Blog  Blog
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

// Blog Blog用設定データ格納構造体
type Blog struct {
	MaxBlogQuota int
}

// GetConfigData 設定データを取得する
func GetConfigData() ConfigData {
	return getConfigData("config.toml")
}

func getConfigData(configFile string) ConfigData {
	var configData ConfigData
	_, err := toml.DecodeFile(configFile, &configData)
	if err != nil {
		panic("tomlファイルを読み込めません")
	}

	// 設定値がない場合、デフォルト値を2019年1月現在の2週間に設定
	if configData.Blog.MaxBlogQuota == 0 {
		configData.Blog.MaxBlogQuota = 2
	}

	return configData
}
