package string

import "strconv"

type Integer struct {
	Builder
	Value int64
}

func Int(i int64) *Integer {
	return &Integer{
		Builder{Base: DefaultStringBase},
		i,
	}
}

func (b *Integer)ToString() string {
	return strconv.FormatInt(b.Value, b.Base)
}
