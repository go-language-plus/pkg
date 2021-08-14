package conc

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {
	type args struct {
		fn func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test_fine",
			args{
				func() {
					fmt.Println("fine")
				},
			},
		},
		{
			"test_panic",
			args{
				func() {
					panic("hehe...")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Go(tt.args.fn)
		})
	}
}
