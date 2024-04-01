package main

import (
	"fmt"
)

func main() {



  fruits := [3]string{"apple", "orange", "banana"}

  for i := 0; i < len(fruits); i++ {
    fmt.Printf("%v\t%v\n",i ,fruits[i]);
	}

  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }

}
