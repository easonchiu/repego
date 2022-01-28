package repego

import (
	"time"
)

// struct
type R struct {
	Count    int
	callback func(r *R) bool
	closed   bool
}

func Call(callback func(r *R) bool) *R {
	return &R{
		Count:    0,
		callback: callback,
		closed:   false,
	}
}

func (r *R) Do(sleep ...time.Duration) bool {
	// initial sleep duration
	var sleepDuration time.Duration = 0
	if sleep != nil && len(sleep) > 0 {
		sleepDuration = sleep[0]
	}

	// call
	for !r.closed {

		r.Count++

		if r.callback(r) {
			r.closed = true
			return true
		}

		// if duration not 0, sleep
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}

	return false
}

func (r *R) Done() bool {
	r.closed = true
	return r.closed
}
