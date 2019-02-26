package config

import "os"

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
	return getConfigData()
}

func getConfigData() ConfigData {
	// 後々awsConfigから取得するように変更したいが、一旦は環境変数から取得する
	slack := Slack{
		SendAPIURL:  os.Getenv("WBEW_SLACK_API_URL"),
		ChannelName: os.Getenv("WBEW_SLACK_CHANNEL_NAME"),
	}
	aws := AWS{
		AccessKey: os.Getenv("WBEW_AWS_ACCESS_KEY"),
		SecretKey: os.Getenv("WBEW_AWS_SECRET_KEY"),
		Region:    os.Getenv("WBEW_DATABASE_REGION"),
		DataBase:  os.Getenv("WBEW_DATABASE_NAME"),
	}
	blog := Blog{
		// デフォルト値を2019年1月現在の2週間に設定
		MaxBlogQuota: 2,
	}
	return ConfigData{
		Slack: slack,
		AWS:   aws,
		Blog:  blog,
	}
}
