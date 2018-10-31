package date

import (
	"time"
)

/**
 * 曜日の番号を返す
 * Goの場合
 * 0 => 日曜
 * 6 => 土曜
 *
 * になるので、Pythonにあわすため、以下にする
 * 0 => 月曜
 * 6 => 日曜
 */
func GetWeekDayNumber() int {
	weekday := int(time.Now().Weekday()) - 1
	if weekday == -1 {
		weekday = 6
	}

	return weekday
}

/**
 * 今週の月曜の日付を取得する
 */
func GetThisMonday(targetHour int) time.Time {
	nowDate := getNowDate(targetHour)
	weekday := GetWeekDayNumber()
	nowDate = time.Date(nowDate.Year(), nowDate.Month(), nowDate.Day(), 00, 00, 00, 0, time.Local)
	return nowDate.Add(time.Duration(-24*weekday) * time.Hour)
}

/**
 * 1週間前の月曜日を取得する
 * (ロジック的には月曜日固定ではないけど、lambdaが月曜日に実行されるからよしとする)
 */
func GetLastWeekMonday(targetHour int) time.Time {
	nowDate := getNowDate(targetHour)
	weekday := 7
	nowDate = time.Date(nowDate.Year(), nowDate.Month(), nowDate.Day(), 00, 00, 00, 0, time.Local)
	return nowDate.Add(time.Duration(-24*weekday) * time.Hour)
}

/**
 * 現在の日付を取得する
 */
func getNowDate(targetHour int) time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), targetHour, 00, 00, 0, time.Local)
}
