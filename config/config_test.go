package config

import (
	"reflect"
	"testing"
)

func Test_getConfigData(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name string
		args args
		want ConfigData
	}{
		{
			name: "normalTest",
			args: args{
				configFile: "normalTest.tomltest",
			},
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
		{
			name: "noBlogConfig",
			args: args{
				configFile: "noBlogConfig.tomltest",
			},
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfigData(tt.args.configFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfigData() = %v, want %v", got, tt.want)
			}
		})
	}
}
