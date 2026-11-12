package main

import (
  "fmt"
)

type kvmem interface {
  Set(k any, v any)
  Get(k any)
  Delete(k any)
}

func main() {
  fmt.Println("I believe this project will be wonderful")
}
