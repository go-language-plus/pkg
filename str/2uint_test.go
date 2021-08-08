package str

import "testing"

func TestStr_Uint(t *testing.T) {
	type fields struct {
		Options Options
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint
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
				Value:   "18446744073709551615",
			},
			18446744073709551615,
			false,
		},
		{
			"testOverflows",
			fields{
				Options: Options{base: DefaultIntBase},
				Value:   "18446744073709551616",
			},
			18446744073709551615,
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
			got, err := s.Uint()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustUint(); got != tt.want {
				t.Errorf("MustUint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Uint16(t *testing.T) {
	type fields struct {
		Options Options
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint16
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
			got, err := s.Uint16()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint16() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustUint16(); got != tt.want {
				t.Errorf("MustUint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Uint32(t *testing.T) {
	type fields struct {
		Options Options
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint32
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
				Value:   "-2147483648",
			},
			0,
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
			got, err := s.Uint32()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustUint32(); got != tt.want {
				t.Errorf("MustUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Uint64(t *testing.T) {
	type fields struct {
		Options Options
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
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
			got, err := s.Uint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustUint64(); got != tt.want {
				t.Errorf("MustUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_Uint8(t *testing.T) {
	type fields struct {
		Options Options
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint8
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
			got, err := s.Uint8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint8() got = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := s.MustUint8(); got != tt.want {
				t.Errorf("MustUint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}
