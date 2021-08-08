package str_test

import (
	"fmt"

	"github.com/go-language-plus/pkg/str"
)

func ExampleUinteger() {
	// int to string
	is := str.Uint(123).String()
	fmt.Printf("%T, %v\n", is, is)

	// Similar to integer

	// Output:
	// string, 123
}
