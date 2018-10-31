package message

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Slackへリマインダーを送信する用のメッセージを作成する
 */
func MakeReminderSendText(targetUserList map[string]int) string {
	textData := fmt.Sprintf(`
<!channel>
まだブログを書けていないユーザーがいます！
今週中に書けるようみんなで煽りましょう！
書けていないユーザー
================
%s
`, strings.Join(getReminderReplaceMessageList(targetUserList), "\n"))

	return textData
}

/**
 * Slackへ先週の結果を送信するようのメッセージを作成する
 */
func MakeResultSendText(targetUserList map[string]int) string {
	textData := fmt.Sprintf(`
<!channel>
1週間お疲れ様でした！
今週も頑張ってブログを書きましょう！
先週ブログを書けていない人は今週書くブログ記事が増えていることを確認してください！
================
%s
`, strings.Join(getReminderReplaceMessageList(targetUserList), "\n"))

	return textData
}

/**
 * リマインダー用の置換文字列リストを取得する
 */
func getReminderReplaceMessageList(targetUserList map[string]int) []string {
	var results []string
	for key, val := range targetUserList {
		results = append(results, "<@"+key+">さん 残り"+strconv.Itoa(val)+"記事")
	}

	return results
}
