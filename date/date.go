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
func GetThisMonday(targetHour int) time.Time {
	nowDate := getNowDate(targetHour)
	weekday := GetWeekDayNumber()
	nowDate = time.Date(nowDate.Year(), nowDate.Month(), nowDate.Day(), 00, 00, 00, 0, time.Local)
	return nowDate.Add(time.Duration(-24*weekday) * time.Hour)
}

// GetLastWeekMonday 1週間前の月曜日を取得する
// GetLastWeekMonday (ロジック的には月曜日固定ではないけど、lambdaが月曜日に実行されるからよしとする)
func GetLastWeekMonday(targetHour int) time.Time {
	nowDate := getNowDate(targetHour)
	weekday := 7
	nowDate = time.Date(nowDate.Year(), nowDate.Month(), nowDate.Day(), 00, 00, 00, 0, time.Local)
	return nowDate.Add(time.Duration(-24*weekday) * time.Hour)
}

// Get2WeekAgoDate 2週間前以上の日付を取得する
func Get2WeekAgoDate() time.Time {
	nowDate := getNowDate(0)
	weekday := 14
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
