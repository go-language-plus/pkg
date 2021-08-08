package str

import (
	"unsafe"
)

type ByteString struct {
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
func UnsafeString(s string) *ByteString {
	str := (*byteStringHeader)(unsafe.Pointer(&s))
	ret := ByteString{
		str: str.str,
		len: str.len,
		cap: str.len,
	}
	return &ret
}

// Bytes convert string to []byte
func (b *ByteString) Bytes() []byte {
	return *(*[]byte)(unsafe.Pointer(b))
}

type sliceByte struct {
	data *[]byte
}

// Bytes return a pointer SliceByte struct instance
func Bytes(b []byte) *sliceByte {
	return &sliceByte{
		data: &b,
	}
}

// String convert []byte to string
func (sb *sliceByte) String() string {
	return *(*string)(unsafe.Pointer(sb.data))
}
