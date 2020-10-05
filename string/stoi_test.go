package string

import (
	"reflect"
	"testing"
)

func TestStr_SetBase(t *testing.T) {
	type fields struct {
		Builder Builder
		Value   string
	}
	type args struct {
		base int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Str
	}{
		{
			"setbase16",
			fields{
				Builder: Builder{
					Base: DefaultStringBase,
				},
				Value: "100",
			},
			args{16},
			&Str{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := String(tt.fields.Value)
			tt.want = s
			if got := s.SetBase(tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Str
	}{
		{
			"testStrInit1",
			args{"0"},
			&Str{
				Builder{
					Base: DefaultStringBase,
				},
				"0",
			},
		},
		{
			"testStrInit2",
			args{"123"},
			&Str{
				Builder{
					Base: DefaultStringBase,
				},
				"123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Int(t *testing.T) {
	type fields struct {
		Builder Builder
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			"test1",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "100",
			},
			100,
			false,
		},
		{
			"test2",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "9223372036854775807",
			},
			9223372036854775807,
			false,
		},
		{
			"testOverflows",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "9223372036854775808",
			},
			9223372036854775807,
			true,
		},
		{
			"testBase2",
			fields{
				Builder: Builder{Base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
		{
			"testBase16",
			fields{
				Builder: Builder{Base: 16},
				Value:   "10000",
			},
			65536,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Builder.Base != DefaultStringBase {
			s.SetBase(tt.fields.Builder.Base)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustInt(); got != tt.want {
				t.Errorf("Int() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Int8(t *testing.T) {
	type fields struct {
		Builder Builder
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int8
		wantErr bool
	}{
		{
			"test1",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "100",
			},
			100,
			false,
		},
		{
			"test2",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "127",
			},
			127,
			false,
		},
		{
			"testOverflows",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "128",
			},
			-128,
			false,
		},
		{
			"test3",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "-128",
			},
			-128,
			false,
		},
		{
			"testOverflows2",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "-129",
			},
			127,
			false,
		},
		{
			"testBase2",
			fields{
				Builder: Builder{Base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Builder.Base != DefaultStringBase {
			s.SetBase(tt.fields.Builder.Base)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Int8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int8() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustInt8(); got != tt.want {
				t.Errorf("Int8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Int16(t *testing.T) {
	type fields struct {
		Builder Builder
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int16
		wantErr bool
	}{
		{
			"test1",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "100",
			},
			100,
			false,
		},
		{
			"test2",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "32767",
			},
			32767,
			false,
		},
		{
			"test3",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "32768",
			},
			-32768,
			false,
		},
		{
			"testBase2",
			fields{
				Builder: Builder{Base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
		{
			"testBase16",
			fields{
				Builder: Builder{Base: 16},
				Value:   "5000",
			},
			20480,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Builder.Base != DefaultStringBase {
			s.SetBase(tt.fields.Builder.Base)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Int16()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int16() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustInt16(); got != tt.want {
				t.Errorf("Int16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Int32(t *testing.T) {
	type fields struct {
		Builder Builder
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int32
		wantErr bool
	}{
		{
			"test1",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "0",
			},
			0,
			false,
		},
		{
			"test2",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "2147483647",
			},
			2147483647,
			false,
		},
		{
			"test3",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "2147483648",
			},
			-2147483648,
			false,
		},
		{
			"testBase2",
			fields{
				Builder: Builder{Base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
		{
			"testBase16",
			fields{
				Builder: Builder{Base: 16},
				Value:   "10000",
			},
			65536,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Builder.Base != DefaultStringBase {
			s.SetBase(tt.fields.Builder.Base)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Int32()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int32() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustInt32(); got != tt.want {
				t.Errorf("Int32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Int64(t *testing.T) {
	type fields struct {
		Builder Builder
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{
			"test1",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "0",
			},
			0,
			false,
		},
		{
			"test2",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "9223372036854775807",
			},
			9223372036854775807,
			false,
		},
		{
			"test3",
			fields{
				Builder: Builder{Base: DefaultStringBase},
				Value:   "9223372036854775808",
			},
			9223372036854775807,
			true,
		},
		{
			"testBase2",
			fields{
				Builder: Builder{Base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
		{
			"testBase16",
			fields{
				Builder: Builder{Base: 16},
				Value:   "10000",
			},
			65536,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Builder.Base != DefaultStringBase {
			s.SetBase(tt.fields.Builder.Base)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Int64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustInt64(); got != tt.want {
				t.Errorf("Int64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
