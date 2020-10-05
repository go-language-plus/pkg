# pkg

## Install
```
go get github.com/go-language-plus/pkg
```

## Usage

### String
import string package:
```go
import "github.com/go-language-plus/pkg/string"
```

Some examples are list as follow:
```go
// string to int
s := string.String("123")

// way in err check
resInt, err := s.Int()
if err != nil {
	panic(err)
}
fmt.Printf("%T, %v\n", resInt, resInt) // int, 123

// way without err check
resInt2 := s.MustInt()
fmt.Printf("%T, %v\n", resInt2, resInt2) // int, 123

// int to string
var i int32 = 1000
resStr := string.Int(int64(i)).String()
fmt.Printf("%T, %v\n", resStr, resStr) // string, 1000

// change default base
resI := string.String("123").SetBase(16).MustInt()
fmt.Printf("%T, %v\n", resI, resI) // int, 291

resS := string.Int(int64(i)).SetBase(16).String()
fmt.Printf("%T, %v\n", resS, resS) // string, 3e8
```
