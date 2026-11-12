package main

import (
  "fmt"
  "sync"
)

type kvmem struct {
  mem map[string]int
  mu sync.Mutex
}

func (kv kvmem) Set(k string, v int) {
  kv.mu.Lock()
  kv.mem[k] = v
  kv.mu.Unlock()
}

func (kv kvmem) Get(k string) int {
  kv.mu.Lock()
  d := kv.mem[k]
  kv.mu.Unlock()
  return d
}

func main() {
  fmt.Println("I believe this project will be wonderful")
  var s kvmem
  s.mem = make(map[string]int)
  s.Set("kirill", 10)
  fmt.Println(s.Get("kirill"))
}
