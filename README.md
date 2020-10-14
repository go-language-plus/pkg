# pkg

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-language-plus/pkg)](https://pkg.go.dev/github.com/go-language-plus/pkg)
![Go](https://github.com/go-language-plus/pkg/workflows/Go/badge.svg?branch=main)
![codecov](https://codecov.io/gh/go-language-plus/pkg/branch/main/graph/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-language-plus/pkg)](https://goreportcard.com/report/github.com/go-language-plus/pkg)

Some explorations of the best practices of the Go language, and simple unified packaging.

This package is just to do some research, and then do some simple packaging based on the standard package or some other operations to unify the troublesome things to facilitate calling, thereby reducing the mental burden. (In other words, yes, this is a toy kit.)

## Install
```
go get -u github.com/go-language-plus/pkg
```

## Usage

In order to make some package names not cause confusion and conflict, we add a "p" (can be understood as the abbreviation of "plus").

### Stringp
Import string package:
```go
import "github.com/go-language-plus/pkg/stringp"
```

Some examples are list as follows:
```go
// string to int
s := stringp.String("123")

resInt, err := s.Int() // // way in err check
if err != nil {
    panic(err)
}
fmt.Printf("%T, %v\n", resInt, resInt) // int, 123

resInt2 := s.MustInt() // // way without err check
fmt.Printf("%T, %v\n", resInt2, resInt2) // int, 123

// int to string
var i int32 = 1000
resStr := stringp.Int(int64(i)).String()
fmt.Printf("%T, %v\n", resStr, resStr) // string, 1000
```
The default `base` parameter above is decimal, if you want to modify the `base` number:
```go
// change default base
resI := stringp.String("123").SetBase(16).MustInt()
fmt.Printf("%T, %v\n", resI, resI) // int, 291

resS := stringp.Int(int64(i)).SetBase(16).String()
fmt.Printf("%T, %v\n", resS, resS) // string, 3e8
```
String package encapsulates the conversion between `string` and `[]byte`. **Note**: This conversion is not safe, you must know what you are doing, otherwise only use it when there is only read-only operation. If you don’t understand, don’t use it, just use the standard conversion.
```go
// string to []byte
b := stringp.ByteString("hello").SliceByte()
fmt.Printf("%T, %v\n", b, b) // []uint8, [104 101 108 108 111]

// []byte to string
bs := stringp.SliceByte([]byte{104, 101, 108, 108, 111}).String()
fmt.Printf("%T, %v\n", bs, bs) // string, hello
```

