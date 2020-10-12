package string

import (
	"reflect"
	"testing"
	"unsafe"
)

func Test_byteString_SliceByte(t *testing.T) {
	type fields struct {
		Str unsafe.Pointer
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
			if got := ByteString("value1").SliceByte(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceByte() = %v, want %v", got, tt.want)
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
			b := []byte{118, 97, 108, 117, 101, 49}
			if got := SliceByte(b).String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
