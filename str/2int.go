package str

import (
	"strconv"
)

// Str the base struct of string for type conversion
type Str struct {
	Options
	Value string
}

// Base set Base to cover default Base setting
func (s *Str) Base(base int) *Str {
	s.base = base
	return s
}

// Unsafe convert to unsafe string type byteString
func (s *Str) Unsafe() *ByteString {
	return UnsafeString(s.Value)
}

// String return a pointer Str struct instance
func String(s string) *Str {
	return &Str{
		Options{base: DefaultIntBase},
		s,
	}
}

// Int convert to int
func (s *Str) Int() (int, error) {
	if s.base == DefaultIntBase {
		return strconv.Atoi(s.Value)
	}

	i64, err := s.parseInt(s.Value, 0)
	return int(i64), err
}

// MustInt convert to int without error checking
func (s *Str) MustInt() int {
	i, _ := s.Int()
	return i
}

// Int8 convert to int8
func (s *Str) Int8() (int8, error) {
	if s.base == DefaultIntBase {
		i, err := strconv.Atoi(s.Value)
		return int8(i), err
	}

	i64, err := s.parseInt(s.Value, 8)
	return int8(i64), err
}

// MustInt8 convert to int8 without error checking
func (s *Str) MustInt8() int8 {
	i, _ := s.Int8()
	return i
}

// Int16 convert to int16
func (s *Str) Int16() (int16, error) {
	if s.base == DefaultIntBase {
		i, err := strconv.Atoi(s.Value)
		return int16(i), err
	}

	i64, err := s.parseInt(s.Value, 16)
	return int16(i64), err
}

// MustInt16 convert to int16 without error checking
func (s *Str) MustInt16() int16 {
	i, _ := s.Int16()
	return i
}

// Int32 convert to int32
func (s *Str) Int32() (int32, error) {
	if s.base == DefaultIntBase {
		i, err := strconv.Atoi(s.Value)
		return int32(i), err
	}

	i64, err := s.parseInt(s.Value, 32)
	return int32(i64), err
}

// MustInt32 convert to int32 without error checking
func (s *Str) MustInt32() int32 {
	i, _ := s.Int32()
	return i
}

// Int64 convert to int64
func (s *Str) Int64() (int64, error) {
	if s.base == DefaultIntBase {
		i, err := strconv.Atoi(s.Value)
		return int64(i), err
	}

	return s.parseInt(s.Value, 64)
}

// MustInt64 convert to int64 without error checking
func (s *Str) MustInt64() int64 {
	i, _ := s.Int64()
	return i
}

// parseInt call strconv.ParseInt
func (s *Str) parseInt(str string, bitSize int) (int64, error) {
	return strconv.ParseInt(str, s.base, bitSize)
}
