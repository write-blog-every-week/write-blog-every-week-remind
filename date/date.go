package date

import (
	"time"
)

// GetWeekDayNumber 曜日の番号を返す
func GetWeekDayNumber() int {
	// Goの場合
	// 0 => 日曜
	// 6 => 土曜
	// になるので、Pythonにあわすため、以下にする
	// 0 => 月曜
	// 6 => 日曜
	weekday := int(time.Now().Weekday()) - 1
	if weekday == -1 {
		weekday = 6
	}

	return weekday
}

// GetThisMonday 今週の月曜の日付を取得する
func GetThisMonday() time.Time {
	nowDate := getNowDate()
	weekday := GetWeekDayNumber()
	return nowDate.AddDate(0, 0, -weekday)
}

// GetLastWeekMonday 1週間前の月曜日を取得する(ロジック的には月曜日固定ではないけど、lambdaが月曜日に実行されるからよしとする)
func GetLastWeekMonday() time.Time {
	nowDate := getNowDate()
	weekday := 7
	return nowDate.AddDate(0, 0, -weekday)
}

/**
 * 現在の日付を取得する
 */
func getNowDate() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 00, 00, 00, 0, time.Local)
}
