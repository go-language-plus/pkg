package string

import (
	"strconv"
	"testing"
)

var s = "2147483647"

// BenchmarkAtoi strconv.Atoi
func BenchmarkAtoi(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_, err := strconv.Atoi(s)
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

// BenchmarkStrconvParseInt strconv.ParseInt
func BenchmarkStrconvParseInt(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

// BenchmarkStringToInt String().ToInt(s)
func BenchmarkStringToInt(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_, err :=  String(s).ToInt()
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}
