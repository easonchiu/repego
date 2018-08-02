package repego

import (
  "reflect"
  "time"
)

// struct
type R struct {
  Count    int
  callback func(r *R)
  close    bool
}

func Call(callback func(r *R)) *R {
  typ := reflect.TypeOf(callback)
  if typ.Kind() != reflect.Func {
    panic("only function can be repeat.")
  }
  return &R{
    Count:    0,
    callback: callback,
    close:    false,
  }
}

func (r *R) Do(sleep ... time.Duration) {
  // initial sleep duration
  var sleepDuration time.Duration = 0
  if sleep != nil && len(sleep) > 0 {
    sleepDuration = sleep[0]
  }

  // call
  for !r.close {

    r.Count ++
    r.callback(r)

    // if duration not 0, sleep
    if sleepDuration > 0 {
      time.Sleep(sleepDuration)
    }
  }
}

func (r *R) Done() {
  r.close = true
}
