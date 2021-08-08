package str_test

import (
	"fmt"

	"github.com/go-language-plus/pkg/str"
)

func ExampleInteger() {
	// int to string
	is := str.Int(123).String()
	fmt.Printf("%T, %v\n", is, is)

	// int8 to string
	var i8 int8 = 127
	i8s := str.Int8(i8).String()
	fmt.Printf("%T, %v\n", i8s, i8s)

	var i16 int16 = 1000
	i16s := str.Int16(i16).String()
	fmt.Printf("%T, %v\n", i16s, i16s)

	// int32 to string
	var i32 int32 = 1000
	i32s := str.Int32(i32).String()
	fmt.Printf("%T, %v\n", i32s, i32s)

	// int64 to string
	var i64 int64 = 1000
	i64s := str.Int64(i64).String()
	fmt.Printf("%T, %v\n", i64s, i64s)

	// All declare function defalt decimal base when convert to string
	// use Base() to set other base
	hex := 0x5AF3ACA48850 // hexadecimal
	hexs := str.Int(hex).Base(16).String()
	fmt.Printf("%T, %v\n", hexs, hexs)

	bin := 0b1111101000 // Binary
	bins := str.Int(bin).Base(2).String()
	fmt.Printf("%T, %v\n", bins, bins)

	// Output:
	// string, 123
	// string, 127
	// string, 1000
	// string, 1000
	// string, 1000
	// string, 5af3aca48850
	// string, 1111101000
}
