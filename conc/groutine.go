package conc

import (
	"fmt"
)

// Go goroutine with defer recover
func Go(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %+v", err)
			}
		}()

		fn()
	}()
}
