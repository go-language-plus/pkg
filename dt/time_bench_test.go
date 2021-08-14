package dt

import (
	"testing"
	"time"
)

/**
goos: darwin
goarch: amd64
pkg: github.com/go-language-plus/pkg/dt
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkTime_NowFormat
BenchmarkTime_NowFormat-16         	 3397639	       321.1 ns/op
BenchmarkTime_DateFormat
BenchmarkTime_DateFormat-16        	   49036	     22779 ns/op
BenchmarkDateTime_NowFormat
BenchmarkDateTime_NowFormat-16     	 2823162	       390.8 ns/op
BenchmarkDateTime_DateFormat
BenchmarkDateTime_DateFormat-16    	   56176	     22071 ns/op
PASS
*/

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

func BenchmarkDateTime_NowFormat(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		Now().Location("UTC").Layout(time.RFC1123).Format()
	}
	b.StopTimer()
}

func BenchmarkDateTime_DateFormat(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		Date(2009, time.November, 10, 23, 0, 0, 0, "Asia/Shanghai").Location(LocationUTC).Layout(time.RFC1123).Format()
	}
	b.StopTimer()
}
