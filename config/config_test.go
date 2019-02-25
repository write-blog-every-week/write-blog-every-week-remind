package config

import (
	"os"
	"reflect"
	"testing"
)

func Test_getConfigData(t *testing.T) {
	tests := []struct {
		name string
		want ConfigData
	}{
		{
			name: "normalTest",
			want: ConfigData{
				Slack: Slack{
					SendAPIURL:  "send_api_url",
					ChannelName: "channel_name",
				},
				AWS: AWS{
					AccessKey: "access_key",
					SecretKey: "secret_key",
					Region:    "region",
					DataBase:  "data_base",
				},
				Blog: Blog{
					MaxBlogQuota: 2,
				},
			},
		},
	}

	resetFunc := []func(){}
	resetFunc = append(resetFunc, setTestEnv("SLACK_API_URL", "send_api_url"))
	resetFunc = append(resetFunc, setTestEnv("SLACK_CHANNEL_NAME", "channel_name"))
	resetFunc = append(resetFunc, setTestEnv("AWS_ACCESS_KEY", "access_key"))
	resetFunc = append(resetFunc, setTestEnv("AWS_SECRET_KEY", "secret_key"))
	resetFunc = append(resetFunc, setTestEnv("DATABASE_REGION", "region"))
	resetFunc = append(resetFunc, setTestEnv("DATABASE_NAME", "data_base"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfigData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfigData() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, f := range resetFunc {
		f()
	}
}

func setTestEnv(key string, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}
