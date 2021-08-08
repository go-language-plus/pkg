package str_test

import (
	"strconv"
	"testing"

	"github.com/go-language-plus/pkg/str"
)

/**
go test -bench=. -run=none -benchmem
=========================================================================================================
goos: darwin
goarch: amd64
pkg: github.com/go-language-plus/pkg/str
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Atoi
Benchmark_Atoi-16                    	137477450	         8.667 ns/op
Benchmark_StrconvParseInt
Benchmark_StrconvParseInt-16         	55469253	        20.24 ns/op
Benchmark_StringToInt
Benchmark_StringToInt-16             	94138113	        11.83 ns/op
Benchmark_StringToSliceByte
Benchmark_StringToSliceByte-16       	229599321	         5.144 ns/op
Benchmark_StringToBytesUnsafe
Benchmark_StringToBytesUnsafe-16     	1000000000	         0.2407 ns/op
Benchmark_BytesToString
Benchmark_BytesToString-16          	330515110	         3.518 ns/op
Benchmark_BytesToStringUnsafe
Benchmark_BytesToStringUnsafe-16    	1000000000	         0.4659 ns/op
PASS
=========================================================================================================
*/

var s = "2147483647"

// Benchmark_Atoi strconv.Atoi
func Benchmark_Atoi(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_, err := strconv.Atoi(s)
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

// Benchmark_StrconvParseInt strconv.ParseInt
func Benchmark_StrconvParseInt(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

// Benchmark_StringToInt String().ToInt(s)
func Benchmark_StringToInt(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_, err := str.String(s).Int()
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func Benchmark_StringToSliceByte(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_ = []byte(s)
	}
	b.StopTimer()
}

func Benchmark_StringToBytesUnsafe(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_ = str.UnsafeString(s).Bytes()
	}
	b.StopTimer()
}

var bs = []byte{118, 97, 108, 117, 101, 49}

func Benchmark_BytesToString(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_ = string(bs)
	}
	b.StopTimer()
}

func Benchmark_BytesToStringUnsafe(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_ = str.Bytes(bs).String()
	}
	b.StopTimer()
}
