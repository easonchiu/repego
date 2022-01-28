package repego

import (
	"time"
)

// struct
type R struct {
	count    int
	maxCount int
	callback func(r *R) bool
	done     bool
}

func Call(callback func(r *R) bool) *R {
	return &R{
		count:    0,
		maxCount: -1,
		callback: callback,
		done:     false,
	}
}

func (r *R) MaxCount(count int) *R {
	if count == 0 {
		panic("the max count can not set to 0.")
	}

	if count < 0 {
		r.maxCount = -1
	} else {
		r.maxCount = count
	}

	return r
}

func (r *R) Do(sleep ...time.Duration) bool {
	// initial sleep duration
	var sleepDuration time.Duration = 0
	if sleep != nil && len(sleep) > 0 {
		sleepDuration = sleep[0]
	}

	// call
	for !r.done && (r.maxCount == -1 || r.count < r.maxCount) {

		r.count++

		if r.callback(r) {
			r.done = true
			return true
		}

		// if duration not 0, sleep
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}

	return false
}

func (r *R) Redo(sleep ...time.Duration) bool {
	if !r.done {
		panic("can not redo now, you must stop it before.")
	}

	r.count = 0
	r.done = false
	return r.Do(sleep...)
}

func (r *R) Done() bool {
	r.done = true
	return r.done
}

func (r *R) IsDone() bool {
	return r.done
}

func (r *R) Count() int {
	return r.count
}
