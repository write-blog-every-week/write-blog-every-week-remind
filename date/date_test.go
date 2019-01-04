package date

import (
	"strconv"
	"testing"
	"time"
)

func TestGetWeekDayNumber(t *testing.T) {
	SetFakeTime(time.Date(2018, 12, 27, 0, 0, 0, 0, asiaTokyo))
	result := GetWeekDayNumber()
	if result != 3 {
		t.Errorf("2018/12/27が木曜日(3)になっていない => number: %s", strconv.Itoa(result))
	}
}

func TestGetThisMonday(t *testing.T) {
	SetFakeTime(time.Date(2018, 12, 27, 0, 0, 0, 0, asiaTokyo))
	result := GetThisMonday()
	if !result.Equal(time.Date(2018, 12, 24, 0, 0, 0, 0, asiaTokyo)) {
		t.Errorf("2018/12/27の週の月曜が2018/12/24ではない => date: %s", result)
	}
}

func TestGetLastWeekMonday(t *testing.T) {
	SetFakeTime(time.Date(2018, 12, 31, 0, 0, 0, 0, asiaTokyo))
	result := GetLastWeekMonday()
	if !result.Equal(time.Date(2018, 12, 24, 0, 0, 0, 0, asiaTokyo)) {
		t.Errorf("2018/12/31に実行した場合の先週の月曜が2018/12/24ではない => date: %s", result)
	}
}
