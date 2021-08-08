package str

import (
	"reflect"
	"testing"
)

func TestUint(t *testing.T) {
	type args struct {
		i uint
	}
	tests := []struct {
		name string
		args args
		want *Uinteger
	}{
		{
			"testuint",
			args{123},
			&Uinteger{
				Options: Options{base: DefaultIntBase},
				Value:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	type args struct {
		i uint16
	}
	tests := []struct {
		name string
		args args
		want *Uinteger
	}{
		{
			"testuint16",
			args{123},
			&Uinteger{
				Options: Options{base: DefaultIntBase},
				Value:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		args args
		want *Uinteger
	}{
		{
			"testuint32",
			args{123},
			&Uinteger{
				Options: Options{base: DefaultIntBase},
				Value:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name string
		args args
		want *Uinteger
	}{
		{
			"testuint64",
			args{123},
			&Uinteger{
				Options: Options{base: DefaultIntBase},
				Value:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		args args
		want *Uinteger
	}{
		{
			"testuint8",
			args{123},
			&Uinteger{
				Options: Options{base: DefaultIntBase},
				Value:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger_Base(t *testing.T) {
	type fields struct {
		Options Options
		Value   uint64
	}
	type args struct {
		base int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Uinteger
	}{
		{
			"setbase16",
			fields{
				Options: Options{
					base: DefaultIntBase,
				},
				Value: 100, // random value
			},
			args{16},
			&Uinteger{
				Options: Options{
					base: 16,
				},
				Value: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Uinteger{
				Options: tt.fields.Options,
				Value:   tt.fields.Value,
			}
			if got := i.Base(tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUinteger_String(t *testing.T) {
	type fields struct {
		Options Options
		Value   uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"test2str",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   100,
			},
			"100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Uinteger{
				Options: tt.fields.Options,
				Value:   tt.fields.Value,
			}
			if got := i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
