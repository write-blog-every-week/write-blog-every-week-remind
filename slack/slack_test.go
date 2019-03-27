package slack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/write-blog-every-week/write-blog-every-week-remind/config"
)

func TestParseSlackParams(t *testing.T) {
	tests := []struct {
		name       string
		rawParams  interface{}
		wantResult *SlackParams
		wantErr    bool
	}{
		{
			name: "nobody",
			rawParams: map[string]interface{}{
				"key": "value",
			},
			wantResult: nil,
			wantErr:    true,
		},
		{
			name: "invalidUrlEscape",
			rawParams: map[string]interface{}{
				"body": "%",
			},
			wantResult: nil,
			wantErr:    true,
		},
		{
			name: "registerBlog",
			rawParams: map[string]interface{}{
				"body": "token=mytoken&user_id=user1&user_name=myuser&text=<https://myawesome.blog.com>",
			},
			wantResult: &SlackParams{
				Token:    "mytoken",
				UserID:   "user1",
				UserName: "myuser",
				Text:     "https://myawesome.blog.com",
			},
			wantErr: false,
		},
		{
			name: "deleteBlogWithAtSign",
			rawParams: map[string]interface{}{
				"body": "token=mytoken&user_id=user1&user_name=myuser&text=@targetuser",
			},
			wantResult: &SlackParams{
				Token:    "mytoken",
				UserID:   "user1",
				UserName: "myuser",
				Text:     "targetuser",
			},
			wantErr: false,
		},
		{
			name: "deleteBlogWithoutAtSign",
			rawParams: map[string]interface{}{
				"body": "token=mytoken&user_id=user1&user_name=myuser&text=targetuser",
			},
			wantResult: &SlackParams{
				Token:    "mytoken",
				UserID:   "user1",
				UserName: "myuser",
				Text:     "targetuser",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ParseSlackParams(tt.rawParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSlackParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ParseSlackParams() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSendMessage(t *testing.T) {
	got := make(map[string]string)
	h := func(w http.ResponseWriter, r *http.Request) {
		bodyStr, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Unexpected error = %v", err)
		}
		var body slackBody
		if err := json.Unmarshal(bodyStr, &body); err != nil {
			t.Errorf("Unexpected error = %v", err)
		}
		got[body.Channel] = body.Text
		w.WriteHeader(http.StatusOK)
	}
	testServer := httptest.NewServer(http.HandlerFunc(h))
	defer testServer.Close()

	type args struct {
		configData config.ConfigData
		sendText   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				configData: config.ConfigData{
					Slack: config.Slack{
						SendAPIURL:  testServer.URL,
						ChannelName: "c1",
					},
					AWS:  config.AWS{},
					Blog: config.Blog{},
				},
				sendText: "hello slack",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendMessage(tt.args.configData, tt.args.sendText)
			if got[tt.args.configData.Slack.ChannelName] != tt.args.sendText {
				t.Errorf("Got = %v, wantChannel = %s, wantSendText = %s", got, tt.args.configData.Slack.ChannelName, tt.args.sendText)
			}
		})
	}
}
