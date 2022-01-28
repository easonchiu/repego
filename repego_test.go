package repego

import (
	"fmt"
	"testing"
)

func TestRepego(t *testing.T) {
	call := Call(func(r *R) bool {
		if r.Count > 10 {
			return r.Done()
		}

		fmt.Println("test:", r.Count)
		return false
	})

	call.Do()
}
