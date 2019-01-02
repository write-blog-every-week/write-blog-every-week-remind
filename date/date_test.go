package date

import (
	"testing"
	"time"
)

func TestGetWeekDayNumber(t *testing.T) {
	tests := []struct {
		name string
		fake time.Time
		want int
	}{
		{
			name: "Thursday 27th Dec 2018",
			fake: time.Date(2018, 12, 27, 0, 0, 0, 0, asiaTokyo),
			want: 3,
		},
		{
			name: "Sunday 23rd Dec 2018",
			fake: time.Date(2018, 12, 23, 0, 0, 0, 0, asiaTokyo),
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetFakeTime(tt.fake)
			if got := GetWeekDayNumber(); got != tt.want {
				t.Errorf("want %d but got %d", tt.want, got)
			}
		})
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
	tests := []struct {
		name string
		fake time.Time
		want time.Time
	}{
		{
			name: "on 2018/12/31 00:00:01 returns 2018/12/24",
			fake: time.Date(2018, 12, 31, 0, 0, 1, 0, asiaTokyo),
			want: time.Date(2018, 12, 24, 0, 0, 0, 0, asiaTokyo),
		},
		{
			name: "on 2018/12/31 00:00:00 returns 2018/12/24",
			fake: time.Date(2018, 12, 31, 0, 0, 0, 0, asiaTokyo),
			want: time.Date(2018, 12, 24, 0, 0, 0, 0, asiaTokyo),
		},
		{
			name: "on 2018/12/30 23:59:59 returns 2018/12/17",
			fake: time.Date(2018, 12, 30, 23, 59, 59, 0, asiaTokyo),
			want: time.Date(2018, 12, 17, 0, 0, 0, 0, asiaTokyo),
		},
		{
			name: "on 2018/12/30 00:00:00 returns 2018/12/17",
			fake: time.Date(2018, 12, 30, 0, 0, 0, 0, asiaTokyo),
			want: time.Date(2018, 12, 17, 0, 0, 0, 0, asiaTokyo),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetFakeTime(tt.fake)
			if got := GetLastWeekMonday(); !got.Equal(tt.want) {
				t.Errorf("want %s but got %s", tt.want, got)
			}
		})
	}
}

func TestSetFakeTime(t *testing.T) {
	// fakeTime初期化
	utc, _ := time.LoadLocation("UTC")
	SetFakeTime(time.Date(1, 1, 1, 0, 0,0, 0, utc))

	tests := []struct {
		name string
		call bool
		fake time.Time
		want time.Time
	}{
		{
			name: "fake time not set",
			call: false,
		},
		{
			name: "fake time set",
			call: true,
			fake: time.Date(2018, 12, 24, 0, 0, 0, 0, asiaTokyo),
			want: time.Date(2018, 12, 24, 0, 0, 0, 0, asiaTokyo),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.call {
				SetFakeTime(tt.fake)
				if fakeTime.IsZero() {
					t.Error("want non-zero fakeTime but got zero fakeTime")
				}
				if !fakeTime.Equal(tt.want) {
					t.Errorf("want %s \nbut fakeTime is %s", tt.want, fakeTime)
				}
			} else {
				if !fakeTime.IsZero() {
					t.Errorf("want zero fakeTime but got %s", fakeTime)
				}
			}
		})
	}
}
