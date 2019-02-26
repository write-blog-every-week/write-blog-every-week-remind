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

	envs := map[string]string{
		"WBEW_SLACK_API_URL":      "send_api_url",
		"WBEW_SLACK_CHANNEL_NAME": "channel_name",
		"WBEW_AWS_ACCESS_KEY":     "access_key",
		"WBEW_AWS_SECRET_KEY":     "secret_key",
		"WBEW_DATABASE_REGION":    "region",
		"WBEW_DATABASE_NAME":      "data_base",
	}
	defer setup(envs)()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfigData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfigData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setup(envs map[string]string) func() {
	pre := map[string]string{}
	for k, v := range envs {
		pre[k] = os.Getenv(k)
		os.Setenv(k, v)
	}
	return func() {
		for k, v := range pre {
			os.Setenv(k, v)
		}
	}
}
