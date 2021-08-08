package str

import "strconv"

// Uinteger the base struct of unsigned integer for type conversion
type Uinteger struct {
	Options
	Value uint64
}

// Base set Base to cover default Base setting
func (i *Uinteger) Base(base int) *Uinteger {
	i.base = base
	return i
}

// Uint64 return a pointer uint64 Uinteger struct instance
func Uint64(i uint64) *Uinteger {
	return &Uinteger{
		Options{base: DefaultIntBase},
		i,
	}
}

// Uint32 return a pointer uint32 Uinteger struct instance
func Uint32(i uint32) *Uinteger {
	return Uint64(uint64(i))
}

// Uint16 return a pointer uint16 Uinteger struct instance
func Uint16(i uint16) *Uinteger {
	return Uint64(uint64(i))
}

// Uint8 return a pointer uint8 Uinteger struct instance
func Uint8(i uint8) *Uinteger {
	return Uint64(uint64(i))
}

// Uint return a pointer uint Uinteger struct instance
func Uint(i uint) *Uinteger {
	return Uint64(uint64(i))
}

// String convert to string
func (i *Uinteger) String() string {
	return strconv.FormatUint(i.Value, i.base)
}
