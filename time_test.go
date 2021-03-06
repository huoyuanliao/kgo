package kgo

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ti := fmt.Sprintf("%d", KTime.Time())
	if len(ti) != 10 {
		t.Error("Time fail")
		return
	}
}

func BenchmarkTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Time()
	}
}

func TestMilliTime(t *testing.T) {
	ti := fmt.Sprintf("%d", KTime.MilliTime())
	if len(ti) != 13 {
		t.Error("MilliTime fail")
		return
	}
}

func BenchmarkMilliTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.MilliTime()
	}
}

func TestMicroTime(t *testing.T) {
	ti := fmt.Sprintf("%d", KTime.MicroTime())
	if len(ti) != 16 {
		t.Error("MicroTime fail")
		return
	}
}

func BenchmarkMicroTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.MicroTime()
	}
}

func TestStr2Timestamp(t *testing.T) {
	ti, err := KTime.Str2Timestamp("2019-07-11 10:11:23")
	if err != nil || ti <= 0 {
		t.Error("Str2Timestamp fail")
		return
	}

	_, err = KTime.Str2Timestamp("02/01/2016 15:04:05")
	if err == nil {
		t.Error("Str2Timestamp fail")
		return
	}

	_, err = KTime.Str2Timestamp("2020-02-01 13:39:36", "2019-07- 11 10: 11:23")
	if err == nil {
		t.Error("Str2Timestamp fail")
		return
	}
}

func BenchmarkStr2Timestamp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KTime.Str2Timestamp("2019-07-11 10:11:23")
	}
}

func TestDate(t *testing.T) {
	date1 := KTime.Date("Y-m-d H:i:s", 1562811851)
	date2 := KTime.Date("y-n-j H:i:s", int64(1562811851))
	date3 := KTime.Date("m/d/y h-i-s", time.Now())
	if date1 == "" || date2 == "" || date3 == "" {
		t.Error("Date fail")
		return
	}

	date4 := KTime.Date("Y-m-d H:i:s")
	date5 := KTime.Date("Y-m-d H:i:s", "hello")
	if date4 == "" || date5 != "" {
		t.Error("Date fail")
		return
	}
}

func BenchmarkDate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Date("Y-m-d H:i:s", 1562811851)
	}
}

func TestCheckDate(t *testing.T) {
	chk1 := KTime.CheckDate(7, 31, 2019)
	chk2 := KTime.CheckDate(2, 31, 2019)
	if !chk1 || chk2 {
		t.Error("CheckDate fail")
		return
	}
	KTime.CheckDate(0, 31, 2019)
	KTime.CheckDate(4, 31, 2019)
	KTime.CheckDate(2, 30, 2008)
}

func BenchmarkCheckDate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.CheckDate(7, 31, 2019)
	}
}

func TestSleep(t *testing.T) {
	ti1 := KTime.Time()
	KTime.Sleep(1)
	ti2 := KTime.Time()
	diff := ti2 - ti1
	if diff != 1 {
		t.Error("Sleep fail")
		return
	}
}

func TestUsleep(t *testing.T) {
	ti1 := KTime.MicroTime()
	KTime.Usleep(100)
	ti2 := KTime.MicroTime()
	diff := ti2 - ti1
	if diff < 100 {
		t.Error("Usleep fail")
		return
	}
}

func TestServiceStartime(t *testing.T) {
	res := KTime.ServiceStartime()
	if res <= 0 {
		t.Error("ServiceStartime fail")
		return
	}
}

func BenchmarkServiceStartime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.ServiceStartime()
	}
}

func TestServiceUptime(t *testing.T) {
	res := KTime.ServiceUptime()
	if int64(res) <= 0 {
		t.Error("ServiceUptime fail")
		return
	}
}

func BenchmarkServiceUptime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.ServiceUptime()
	}
}

func TestGetMonthDays(t *testing.T) {
	var tests = []struct {
		month    int
		year     int
		expected int
	}{
		{1, 2009, 31},
		{0, 2009, 0},
		{2, 2009, 28},
		{2, 2016, 29},
		{2, 1900, 28},
		{2, 1600, 29},
	}
	for _, test := range tests {
		actual := KTime.GetMonthDays(test.month, test.year)
		if actual != test.expected {
			t.Errorf("Expected GetMonthDays(%d, %d) to be %v, got %v", test.month, test.year, test.expected, actual)
		}
	}

	KTime.GetMonthDays(2)
	KTime.GetMonthDays(3, 1970)
}

func BenchmarkGetMonthDays(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.GetMonthDays(3, 1970)
	}
}

func TestYearMonthDay(t *testing.T) {
	y := KTime.Year()
	m := KTime.Month()
	d := KTime.Day()
	if y <= 0 || m <= 0 || d < 0 {
		t.Error("Year/Month/Day fail")
		return
	}

	tim, _ := KTime.Str2Timestruct("2019-07-11 10:11:23")
	y = KTime.Year(tim)
	m = KTime.Month(tim)
	d = KTime.Day(tim)
	if y != 2019 || m != 7 || d != 11 {
		t.Error("Year/Month/Day fail")
		return
	}
}

func BenchmarkYear(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Year()
	}
}

func BenchmarkMonth(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Month()
	}
}

func BenchmarkDay(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Day()
	}
}

func TestHourMinuteSecond(t *testing.T) {
	h := KTime.Hour()
	m := KTime.Minute()
	s := KTime.Second()
	if h < 0 || m < 0 || s < 0 {
		t.Error("Hour/Minute/Second fail")
		return
	}

	tim, _ := KTime.Str2Timestruct("2019-07-11 10:11:23")
	h = KTime.Hour(tim)
	m = KTime.Minute(tim)
	s = KTime.Second(tim)
	if h != 10 || m != 11 || s != 23 {
		t.Error("Hour/Minute/Second fail")
		return
	}
}

func BenchmarkHour(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Hour()
	}
}

func BenchmarkMinute(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Minute()
	}
}

func BenchmarkSecond(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Second()
	}
}
