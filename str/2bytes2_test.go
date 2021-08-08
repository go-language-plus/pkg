package str

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestByteString_Bytes(t *testing.T) {
	type fields struct {
		str unsafe.Pointer
		len int
		cap int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			"test",
			fields{},
			[]byte{118, 97, 108, 117, 101, 49},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := UnsafeString("value1")
			if got := b.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteString.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceByte_String(t *testing.T) {
	type fields struct {
		data *[]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"test",
			fields{},
			"value1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := Bytes([]byte{118, 97, 108, 117, 101, 49})
			if got := sb.String(); got != tt.want {
				t.Errorf("sliceByte.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
