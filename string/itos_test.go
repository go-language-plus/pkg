package string

import "testing"

func TestInteger_ToString(t *testing.T) {
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
				Builder: Builder{Base: 10},
				Value: 100,
			},
			"100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.fields.Value).ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}