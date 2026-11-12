package main

import (
  "fmt"
  "sync"
)

type kvmem struct {
  mem map[any]any
  mu sync.Mutex
}

func main() {
  fmt.Println("I believe this project will be wonderful")
}
