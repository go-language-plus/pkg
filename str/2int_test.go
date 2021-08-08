package str

import "testing"

func TestStr_Int(t *testing.T) {
	type fields struct {
		Options Options
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
				Options: Options{base: DefaultIntBase},
				Value:   "100",
			},
			100,
			false,
		},
		{
			"test2",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "9223372036854775807",
			},
			9223372036854775807,
			false,
		},
		{
			"testOverflows",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "9223372036854775808",
			},
			9223372036854775807,
			true,
		},
		{
			"testBase2",
			fields{
				Options: Options{base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
		{
			"testBase16",
			fields{
				Options: Options{base: 16},
				Value:   "10000",
			},
			65536,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Options.base != DefaultIntBase {
			s.Base(tt.fields.Options.base)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("Str.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Str.Int() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustInt(); got != tt.want {
				t.Errorf("MustInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Int8(t *testing.T) {
	type fields struct {
		Options Options
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
				Options: Options{base: DefaultIntBase},
				Value:   "100",
			},
			100,
			false,
		},
		{
			"test2",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "127",
			},
			127,
			false,
		},
		{
			"testOverflows",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "128",
			},
			-128,
			false,
		},
		{
			"test3",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "-128",
			},
			-128,
			false,
		},
		{
			"testOverflows2",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "-129",
			},
			127,
			false,
		},
		{
			"testBase2",
			fields{
				Options: Options{base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Options.base != DefaultIntBase {
			s.Base(tt.fields.Options.base)
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
				t.Errorf("MustInt8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Int16(t *testing.T) {
	type fields struct {
		Options Options
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
				Options: Options{base: DefaultIntBase},
				Value:   "100",
			},
			100,
			false,
		},
		{
			"test2",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "32767",
			},
			32767,
			false,
		},
		{
			"test3",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "32768",
			},
			-32768,
			false,
		},
		{
			"testBase2",
			fields{
				Options: Options{base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
		{
			"testBase16",
			fields{
				Options: Options{base: 16},
				Value:   "5000",
			},
			20480,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Options.base != DefaultIntBase {
			s.Base(tt.fields.Options.base)
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
				t.Errorf("MustInt16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Int32(t *testing.T) {
	type fields struct {
		Options Options
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
				Options: Options{base: DefaultIntBase},
				Value:   "0",
			},
			0,
			false,
		},
		{
			"test2",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "2147483647",
			},
			2147483647,
			false,
		},
		{
			"test3",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "2147483648",
			},
			-2147483648,
			false,
		},
		{
			"testBase2",
			fields{
				Options: Options{base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
		{
			"testBase16",
			fields{
				Options: Options{base: 16},
				Value:   "10000",
			},
			65536,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Options.base != DefaultIntBase {
			s.Base(tt.fields.Options.base)
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
				t.Errorf("MustInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Int64(t *testing.T) {
	type fields struct {
		Options Options
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
				Options: Options{base: DefaultIntBase},
				Value:   "0",
			},
			0,
			false,
		},
		{
			"test2",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "9223372036854775807",
			},
			9223372036854775807,
			false,
		},
		{
			"test3",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "9223372036854775808",
			},
			9223372036854775807,
			true,
		},
		{
			"testBase2",
			fields{
				Options: Options{base: 2},
				Value:   "10000",
			},
			16,
			false,
		},
		{
			"testBase16",
			fields{
				Options: Options{base: 16},
				Value:   "10000",
			},
			65536,
			false,
		},
	}
	for _, tt := range tests {
		s := String(tt.fields.Value)
		if tt.fields.Options.base != DefaultIntBase {
			s.Base(tt.fields.Options.base)
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
				t.Errorf("MustInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
