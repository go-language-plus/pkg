package string

import "strconv"

type Integer struct {
	Builder
	Value int64
}

// Int return a pointer Integer struct instance
func Int(i int64) *Integer {
	return &Integer{
		Builder{Base: DefaultStringBase},
		i,
	}
}

// SetBase set Base to cover default Base setting
func (i *Integer) SetBase(base int) *Integer {
	i.Base = base
	return i
}

// String convert to string
func (i *Integer) String() string {
	return strconv.FormatInt(i.Value, i.Base)
}
