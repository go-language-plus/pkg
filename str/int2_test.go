package str

import (
	"reflect"
	"testing"
)

func TestInteger_Base(t *testing.T) {
	type fields struct {
		Builder Options
		Value   int
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
				Builder: Options{
					base: DefaultIntBase,
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
			if got := i.Base(tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger_String(t *testing.T) {
	type fields struct {
		Options Options
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
				Options: Options{base: DefaultIntBase},
				Value:   100,
			},
			"100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Integer{
				Options: tt.fields.Options,
				Value:   tt.fields.Value,
			}
			if got := i.String(); got != tt.want {
				t.Errorf("Integer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
