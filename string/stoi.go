package string

import (
	"strconv"
)

type Str struct {
	Builder
	Value string
}

func String(s string) *Str {
	return &Str{
		Builder{Base: DefaultStringBase},
		s,
	}
}

func (s *Str) ToInt() (int, error)  {
	if s.Base == DefaultStringBase {
		return strconv.Atoi(s.Value)
	}

	i64, err := s.parseInt(s.Value, 0)
	return int(i64), err
}

func (s *Str) ToInt8() (int8, error)  {
	if s.Base == DefaultStringBase {
		i, err := strconv.Atoi(s.Value)
		return int8(i), err
	}

	i64, err := s.parseInt(s.Value, 8)
	return int8(i64), err
}

func (s *Str) ToInt32() (int32, error)  {
	if s.Base == DefaultStringBase {
		i, err := strconv.Atoi(s.Value)
		return int32(i), err
	}

	i64, err := s.parseInt(s.Value, 32)
	return int32(i64), err
}

func (s *Str) ToInt64() (int64, error)  {
	if s.Base == DefaultStringBase {
		i, err := strconv.Atoi(s.Value)
		return int64(i), err
	}

	return s.parseInt(s.Value, 64)
}

func (s *Str)parseInt(str string, bitSize int) (int64, error) {
	return strconv.ParseInt(str, s.Base, bitSize)
}
