package string

import (
	"unsafe"
)

type byteString struct {
	str unsafe.Pointer
	len int
	cap int
}

type byteStringHeader struct {
	str unsafe.Pointer
	len int
}

// ByteString return a pointer ByteString struct instance
// Note: Independent ByteString comes out because it is special
func ByteString(s string) *byteString {
	str := (*byteStringHeader)(unsafe.Pointer(&s))
	ret := byteString{
		str: str.str,
		len: str.len,
		cap: str.len,
	}
	return &ret
}

// SliceByte convert string to []byte
func (b *byteString) SliceByte() []byte {
	return *(*[]byte)(unsafe.Pointer(b))
}

type sliceByte struct {
	data *[]byte
}

// SliceByte return a pointer SliceByte struct instance
func SliceByte(b []byte) *sliceByte {
	return &sliceByte{
		data: &b,
	}
}

// String convert []byte to string
func (sb *sliceByte) String() string {
	return *(*string)(unsafe.Pointer(sb.data))
}
