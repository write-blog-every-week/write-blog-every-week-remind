package message

import (
	"testing"

	"github.com/write-blog-every-week/write-blog-every-week-remind/database"
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
書けていないユーザー: 1人
================
<@fuga>さん    残り1記事
`,
		},
		{
			name: "tenUsers",
			list: map[string]int{
				"user1":  1,
				"user2":  1,
				"user3":  1,
				"user4":  1,
				"user5":  2,
				"user6":  2,
				"user7":  2,
				"user8":  2,
				"user9":  2,
				"user10": 2,
			},
			want: `
<!channel>
まだブログを書けていないユーザーがいます！
今週中に書けるようみんなで煽りましょう！
書けていないユーザー: 10人
================
<@user9>さん     残り2記事
<@user8>さん     残り2記事
<@user7>さん     残り2記事
<@user6>さん     残り2記事
<@user5>さん     残り2記事
<@user10>さん    残り2記事
<@user4>さん     残り1記事
<@user3>さん     残り1記事
<@user2>さん     残り1記事
<@user1>さん     残り1記事
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
		name string
		list map[string]int
		want string
	}{
		{
			name: "normalTest",
			list: map[string]int{
				"fuga": 3,
			},
			want: `残念ながら以下の方は退会となります :cry:
================
<@fuga>さん    残り3記事
================`,
		},
		{
			name: "zeroUserGreaterThanQuota",
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
				"hoge":         1,
				"barbar":       1,
				"hogehogehoge": 2,
				"barbarbar":    2,
				"fuga":         2,
			},
			want: `<@hogehogehoge>さん    残り2記事
<@fuga>さん            残り2記事
<@barbarbar>さん       残り2記事
<@hoge>さん            残り1記事
<@barbar>さん          残り1記事
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

func TestCreateFailedRSSMessage(t *testing.T) {
	tests := []struct {
		name string
		list []*database.WriteBlogEveryWeek
		want string
	}{
		{
			name: "multiple",
			list: []*database.WriteBlogEveryWeek{
				&database.WriteBlogEveryWeek{
					UserID:       "hoge",
					UserName:     "budougumi0617",
					FeedURL:      "https://budougumi0617.github.io/index.xml",
					RequireCount: 1,
				},
				&database.WriteBlogEveryWeek{
					UserID:       "hoge",
					UserName:     "budougumi0618",
					FeedURL:      "https://budougumi0618.github.io/index.xml",
					RequireCount: 1,
				},
			},
			want: `以下の方々のRSSの読み込みに失敗しました :scream:
================
<@budougumi0617>:    https://budougumi0617.github.io/index.xml
<@budougumi0618>:    https://budougumi0618.github.io/index.xml
================
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateFailedRSSMessage(tt.list); got != tt.want {
				t.Errorf("want \n%s\n, but got \n%s\n", tt.want, got)
			}
		})
	}
}
