package main

import (
	"fmt"
)

func myFunction(x int, y string) (result int, txt1 string) {
   result = x + x
   txt1 = y + " World!"
  return result, txt1
}

func main() {
  a, b := myFunction(5, "Hello")
  fmt.Println(a, b)
}
