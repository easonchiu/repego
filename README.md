# Repego

重复循环调用函数方法

```
repego.Call(func(r *R) bool {
  fmt.Println("repeat...")
  
  return false
}).MaxCount(100).Do()
```

```
r := repego(func(r *R) bool {
  fmt.Println("repeat...")
  
  return false
})

// start
r.Do()

// check is done
r.IsDone()

// set max count
r.MaxCount(100)

// redo
if r.IsDone() {
  r.MaxCount(20).Redo()
}

// with sleep
r.Do(time.Second * 1)

```