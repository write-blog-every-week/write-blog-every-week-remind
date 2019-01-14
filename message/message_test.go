package message

import (
	"testing"
)

func TestMakeReminderSendText(t *testing.T) {
	tests := []struct {
		name      string
		list      map[string]int
		want      string
		wantError bool
		err       error
	}{
		{
			name: "zero",
			list: map[string]int{},
			want: `
<!channel>
今週は全員がブログを書きました！ :tada:
やったね！！！
`,
		},
		{
			name: "single",
			list: map[string]int{
				"fuga": 1,
			},
			want: `
<!channel>
まだブログを書けていないユーザーがいます！
今週中に書けるようみんなで煽りましょう！
書けていないユーザー
================
<@fuga>さん    残り1記事
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeReminderSendText(tt.list); got != tt.want {
				t.Errorf("want \n%s\n, but got \n%s\n", tt.want, got)
			}
		})
	}
}

func TestMakeResultSendText(t *testing.T) {
	tests := []struct {
		name         string
		maxBlogQuota int
		list         map[string]int
		want         string
	}{
		{
			name:         "normalTest",
			maxBlogQuota: 2,
			list: map[string]int{
				"user1": 2,
				"user2": 1,
			},
			want: `
<!channel>
1週間お疲れ様でした！
今週も頑張ってブログを書きましょう！
先週ブログを書けていない人は今週書くブログ記事が増えていることを確認してください！
================
<@user1>さん    残り2記事
<@user2>さん    残り1記事
================

今週は退会対象者はいません！ :tada:
`,
		},
		{
			name:         "1userOverQuota",
			maxBlogQuota: 2,
			list: map[string]int{
				"user1": 2,
				"user2": 1,
				"user3": 3,
			},
			want: `
<!channel>
1週間お疲れ様でした！
今週も頑張ってブログを書きましょう！
先週ブログを書けていない人は今週書くブログ記事が増えていることを確認してください！
================
<@user1>さん    残り2記事
<@user2>さん    残り1記事
================

残念ながら以下の方は退会となります :cry:
================
<@user3>さん    残り3記事
================
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeResultSendText(tt.maxBlogQuota, tt.list); got != tt.want {
				t.Errorf("want \n%s\n, but got\n%s\n", tt.want, got)
			}
		})
	}
}

func Test_getCancelReplaceMessageList(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name         string
		list         map[string]int
		want         string
	}{
		{
			name:         "normalTest",
			list: map[string]int{
				"fuga": 3,
			},
			want: `残念ながら以下の方は退会となります :cry:
================
<@fuga>さん    残り3記事
================`,
		},
		{
			name:         "zeroUserGreaterThanQuota",
			list: map[string]int{},
			want: "今週は退会対象者はいません！ :tada:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCancelReplaceMessageList(tt.list); got != tt.want {
				t.Errorf("want \n%s\n, but got \n%s\n", tt.want, got)
			}
		})
	}
}

func TestGetRminderReplaceMessageList(t *testing.T) {
	tests := []struct {
		name      string
		list      map[string]int
		want      string
		wantError bool
		err       error
	}{
		{
			name: "single",
			list: map[string]int{
				"hoge": 2,
			},
			want: "<@hoge>さん    残り2記事\n",
		},
		{
			name: "multiple",
			list: map[string]int{
				"hoge":         2,
				"barbar":       30,
				"hogehogehoge": 100000000,
			},
			want: `<@barbar>さん          残り30記事
<@hoge>さん            残り2記事
<@hogehogehoge>さん    残り100000000記事
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getReminderReplaceMessageList(tt.list); got != tt.want {
				t.Errorf("want \n%s\n, but got \n%s\n", tt.want, got)
			}
		})
	}
}
