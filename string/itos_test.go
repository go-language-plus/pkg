package string

import (
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	var testInt = 1

	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want *Integer
	}{
		{
			"testInit",
			args{1},
			&Integer{
				Builder{
					Base: DefaultStringBase,
				},
				1,
			},
		},
		{
			"testInit2",
			args{int64(testInt)},
			&Integer{
				Builder{
					Base: DefaultStringBase,
				},
				int64(testInt),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger_SetBase(t *testing.T) {
	type fields struct {
		Builder Builder
		Value   int64
	}
	type args struct {
		base int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Integer
	}{
		{
			"setbase16",
			fields{
				Builder: Builder{
					Base: DefaultStringBase,
				},
				Value: 100, // random value
			},
			args{16},
			&Integer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Int(tt.fields.Value)
			tt.want = i
			if got := i.SetBase(tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger_String(t *testing.T) {
	type fields struct {
		Builder Builder
		Value   int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"test1",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   100,
			},
			"100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.fields.Value).String(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
