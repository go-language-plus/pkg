package string

import (
	"strconv"
	"testing"
)

// go test -bench=. -run=none -benchmem
// =========================================================================================================
// goos: darwin
// goarch: amd64
// pkg: github.com/go-language-plus/pkg/string
// BenchmarkAtoi-16                136491384                9.04 ns/op            0 B/op          0 allocs/op
// BenchmarkStrconvParseInt-16     60729482                20.1 ns/op             0 B/op          0 allocs/op
// BenchmarkStringToInt-16         98901199                12.2 ns/op             0 B/op          0 allocs/op
// PASS
// =========================================================================================================

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
		_, err := String(s).Int()
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}
