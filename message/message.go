package message

import (
	"bytes"
	"fmt"
	"sort"
	"text/tabwriter"
)

// MakeReminderSendText Slackへリマインダーを送信する用のメッセージを作成する
func MakeReminderSendText(targetUserList map[string]int) string {
	if len(targetUserList) == 0 {
		return `
<!channel>
今週は全員がブログを書きました！ :tada:
やったね！！！
`
	}
	return fmt.Sprintf(`
<!channel>
まだブログを書けていないユーザーがいます！
今週中に書けるようみんなで煽りましょう！
書けていないユーザー
================
%s`, getReminderReplaceMessageList(targetUserList))
}

// MakeResultSendText Slackへ先週の結果を送信するようのメッセージを作成する
func MakeResultSendText(targetUserList map[string]int) string {
	return fmt.Sprintf(`
<!channel>
1週間お疲れ様でした！
今週も頑張ってブログを書きましょう！
先週ブログを書けていない人は今週書くブログ記事が増えていることを確認してください！
================
%s`, getReminderReplaceMessageList(targetUserList))
}

// getReminderReplaceMessageList リマインダー用のユーザーリスト文字列を取得する
func getReminderReplaceMessageList(targetUserList map[string]int) string {
	var buf bytes.Buffer
	tw := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)
	names := make([]string, 0, len(targetUserList))
	for name := range targetUserList {
		names = append(names, name)
	}
	sort.Strings(names) //sort by key
	for _, n := range names {
		tw.Write([]byte(fmt.Sprintf("<@%s>さん\t残り%d記事\n", n, targetUserList[n])))
	}
	if err := tw.Flush(); err != nil {
		return fmt.Sprintf("リスト生成に失敗 %+v\n", targetUserList)
	}

	return buf.String()
}
