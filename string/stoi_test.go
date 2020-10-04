package string

import (
	"testing"
)

func TestStr_ToInt(t *testing.T) {
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
				Builder: Builder{Base: 10},
				Value:"100",
			},
			100,
			false,
		},
		{
			"test2",
			fields{
				Builder: Builder{Base: 10},
				Value:"129",
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := String(tt.fields.Value).ToInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
