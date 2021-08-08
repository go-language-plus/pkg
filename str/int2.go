package str

import "strconv"

// Integer the base struct of integer for type conversion
type Integer struct {
	Options
	Value int64
}

// Base set Base to cover default Base setting
func (i *Integer) Base(base int) *Integer {
	i.base = base
	return i
}

// Int64 return a pointer int64 Integer struct instance
func Int64(value int64) *Integer {
	return &Integer{
		Options{base: DefaultIntBase},
		value,
	}
}

// Int32 return a pointer int32 Integer struct instance
func Int32(i int32) *Integer {
	return Int64(int64(i))
}

// Int16 return a pointer int16 Integer struct instance
func Int16(i int16) *Integer {
	return Int64(int64(i))
}

// Int8 return a pointer int8 Integer struct instance
func Int8(i int8) *Integer {
	return Int64(int64(i))
}

// Int return a pointer Integer struct instance
func Int(i int) *Integer {
	return Int64(int64(i))
}

// String convert to string
func (i *Integer) String() string {
	return strconv.FormatInt(i.Value, i.base)
}
