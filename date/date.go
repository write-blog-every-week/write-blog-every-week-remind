package date

import (
	"time"
)

var asiaTokyo, _ = time.LoadLocation("Asia/Tokyo")
var fakeTime time.Time

// GetWeekDayNumber 曜日の番号を返す
func GetWeekDayNumber() int {
	// Goの場合
	// 0 => 日曜
	// 6 => 土曜
	// になるので、Pythonにあわすため、以下にする
	// 0 => 月曜
	// 6 => 日曜
	weekday := int(TimeNow().Weekday()) - 1
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

// GetLastWeekMonday 1週間前の月曜日を取得する
func GetLastWeekMonday() time.Time {
	nowDate := getNowDate()
	weekday := 7 + GetWeekDayNumber()
	return nowDate.AddDate(0, 0, -weekday)
}

// SetFakeTime stub用の日付データをセットする
func SetFakeTime(t time.Time) {
	fakeTime = t
}

// TimeNow 現在の日付を返す(stub用のデータが有る場合はstubデータを返す)
func TimeNow() time.Time {
	if !fakeTime.IsZero() {
		return fakeTime
	}
	return time.Now()
}

/**
 * 現在の日付を取得する
 */
func getNowDate() time.Time {
	t := TimeNow()
	return time.Date(t.Year(), t.Month(), t.Day(), 00, 00, 00, 0, asiaTokyo)
}
