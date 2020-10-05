package string

import (
	"strconv"
)

type Str struct {
	Builder
	Value string
}

func (s *Str) SetBase(base int) *Str {
	s.Base = base
	return s
}

func String(s string) *Str {
	return &Str{
		Builder{Base: DefaultStringBase},
		s,
	}
}

func (s *Str) Int() (int, error) {
	if s.Base == DefaultStringBase {
		return strconv.Atoi(s.Value)
	}

	i64, err := s.parseInt(s.Value, 0)
	return int(i64), err
}

func (s *Str) MustInt() int {
	i, _ := s.Int()
	return i
}

func (s *Str) Int8() (int8, error) {
	if s.Base == DefaultStringBase {
		i, err := strconv.Atoi(s.Value)
		return int8(i), err
	}

	i64, err := s.parseInt(s.Value, 8)
	return int8(i64), err
}

func (s *Str) MustInt8() int8 {
	i, _ := s.Int8()
	return i
}

func (s *Str) Int16() (int16, error) {
	if s.Base == DefaultStringBase {
		i, err := strconv.Atoi(s.Value)
		return int16(i), err
	}

	i64, err := s.parseInt(s.Value, 32)
	return int16(i64), err
}

func (s *Str) MustInt16() int16 {
	i, _ := s.Int16()
	return i
}

func (s *Str) Int32() (int32, error) {
	if s.Base == DefaultStringBase {
		i, err := strconv.Atoi(s.Value)
		return int32(i), err
	}

	i64, err := s.parseInt(s.Value, 32)
	return int32(i64), err
}

func (s *Str) MustInt32() int32 {
	i, _ := s.Int32()
	return i
}

func (s *Str) Int64() (int64, error) {
	if s.Base == DefaultStringBase {
		i, err := strconv.Atoi(s.Value)
		return int64(i), err
	}

	return s.parseInt(s.Value, 64)
}

func (s *Str) MustInt64() int64 {
	i, _ := s.Int64()
	return i
}

func (s *Str) parseInt(str string, bitSize int) (int64, error) {
	return strconv.ParseInt(str, s.Base, bitSize)
}
