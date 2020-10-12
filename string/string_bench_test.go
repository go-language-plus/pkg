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
// BenchmarkAtoi-16                        129234172                9.24 ns/op            0 B/op          0 allocs/op
// BenchmarkStrconvParseInt-16             57918210                21.0 ns/op             0 B/op          0 allocs/op
// BenchmarkStringToInt-16                 91119936                12.7 ns/op             0 B/op          0 allocs/op
// BenchmarkStringToSliceByte-16           220481298                5.58 ns/op            0 B/op          0 allocs/op
// BenchmarkStringToSliceByteUnsafe-16     1000000000               0.261 ns/op           0 B/op          0 allocs/op
// BenchmarkSliceByteToString-16           298688124                3.80 ns/op            0 B/op          0 allocs/op
// BenchmarkSliceByteToStringUnsafe-16     1000000000               0.502 ns/op           0 B/op          0 allocs/op
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

func BenchmarkStringToSliceByte(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_ = []byte(s)
	}
	b.StopTimer()
}

func BenchmarkStringToSliceByteUnsafe(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_ = ByteString(s).SliceByte()
	}
	b.StopTimer()
}

var bs = []byte{118, 97, 108, 117, 101, 49}

func BenchmarkSliceByteToString(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_ = string(bs)
	}
	b.StopTimer()
}

func BenchmarkSliceByteToStringUnsafe(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		_ = SliceByte(bs).String()
	}
	b.StopTimer()
}
