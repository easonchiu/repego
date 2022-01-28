package repego

import (
	"fmt"
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	call := Call(func(r *R) bool {
		if r.Count() > 10 {
			return r.Done()
		}

		fmt.Println("test default:", r.Count())
		return false
	})

	call.Do(time.Second)

	fmt.Println("is done:", call.IsDone())
}

func TestMaxCount(t *testing.T) {
	call := Call(func(r *R) bool {
		fmt.Println("test max count:", r.Count())
		return false
	})

	call.MaxCount(50).Do()
}
