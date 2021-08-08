package str_test

import (
	"fmt"

	"github.com/go-language-plus/pkg/str"
)

func ExampleStr() {
	// string to int
	resInt, err := str.String("123").Int() // way in err check
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T, %v\n", resInt, resInt)

	resInt2 := str.String("123").MustInt() // way without err check
	fmt.Printf("%T, %v\n", resInt2, resInt2)

	// change default base
	resI := str.String("5af3aca48850").Base(16).MustInt()
	fmt.Printf("%T, %v\n", resI, resI)

	// can use unsafe string of conert to []byte
	resbt := str.String("123").Unsafe().Bytes()
	fmt.Printf("%T, %v\n", resbt, resbt)

	// Output:
	// int, 123
	// int, 123
	// int, 100002620016720
	// []uint8, [49 50 51]
}
