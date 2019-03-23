package slack

import (
	"reflect"
	"testing"
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
