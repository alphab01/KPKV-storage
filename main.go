package main

import (
  "fmt"
  "sync"
)

type kvmem struct {
  mem map[any]any
  mu sync.Mutex
}

func (kv kvmem) Set(k any, v any) {
  kv.mu.Lock()
  kv.mem[k] = v
  kv.mu.Unlock()
}

func main() {
  fmt.Println("I believe this project will be wonderful")
}
