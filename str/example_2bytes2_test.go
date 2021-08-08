package str_test

import (
	"fmt"

	"github.com/go-language-plus/pkg/str"
)

func ExampleByteString() {
	// string to []byte
	b := str.UnsafeString("hello").Bytes()
	fmt.Printf("%T, %v\n", b, b) // []uint8, [104 101 108 108 111]

	// []byte to string
	bs := str.Bytes([]byte{104, 101, 108, 108, 111}).String()
	fmt.Printf("%T, %v\n", bs, bs) // string, hello

	// Output:
	// []uint8, [104 101 108 108 111]
	// string, hello
}
