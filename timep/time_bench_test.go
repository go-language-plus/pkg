package timep

import (
	"testing"
	"time"
)

func BenchmarkTime_NowFormat(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		time.Now().In(time.UTC).Format(time.RFC1123)
	}
	b.StopTimer()
}

func BenchmarkTime_DateFormat(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		time.Date(2009, time.November, 10, 23, 0, 0, 0, loc).In(time.UTC).Format(time.RFC1123)
	}
	b.StopTimer()
}

func BenchmarkTimeP_NowFormat(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		Now().Location("UTC").Layout(time.RFC1123).Format()
	}
	b.StopTimer()
}

func BenchmarkTimeP_DateFormat(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		Date(2009, time.November, 10, 23, 0, 0, 0, "Asia/Shanghai").Location(LocationUTC).Layout(time.RFC1123).Format()
	}
	b.StopTimer()
}
