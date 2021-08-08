package str

import (
	"strconv"
)

// Uint convert to uint
func (s *Str) Uint() (uint, error) {
	i, err := s.parseUint(s.Value, s.base, 0)
	return uint(i), err
}

// MustUint convert to int without error checking
func (s *Str) MustUint() uint {
	i, _ := s.Uint()
	return i
}

// Uint8 convert to uint8
func (s *Str) Uint8() (uint8, error) {
	i, err := s.parseUint(s.Value, s.base, 8)
	return uint8(i), err
}

// MustUint8 convert to uint8 without error checking
func (s *Str) MustUint8() uint8 {
	i, _ := s.Uint8()
	return i
}

// Uint16 convert to uint16
func (s *Str) Uint16() (uint16, error) {
	i, err := s.parseUint(s.Value, s.base, 16)
	return uint16(i), err
}

// MustUint16 convert to uint16 without error checking
func (s *Str) MustUint16() uint16 {
	i, _ := s.Uint16()
	return i
}

// Uint32 convert to uint32
func (s *Str) Uint32() (uint32, error) {
	i, err := s.parseUint(s.Value, s.base, 32)
	return uint32(i), err
}

// MustUint32 convert to uint32 without error checking
func (s *Str) MustUint32() uint32 {
	i, _ := s.Uint32()
	return i
}

// Uint64 convert to uint64
func (s *Str) Uint64() (uint64, error) {
	return s.parseUint(s.Value, s.base, 64)
}

// MustUint64 convert to int64 without error checking
func (s *Str) MustUint64() uint64 {
	i, _ := s.Uint64()
	return i
}

// parseInt call strconv.ParseInt
func (s *Str) parseUint(str string, base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(str, base, bitSize)
}
