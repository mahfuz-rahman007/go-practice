package main

import (
	"fmt"
)

func main() {

	for i := 10; i > 0; i-- {

		for j := 0; j < i; j++ {
			fmt.Print(j)
		}

		fmt.Print("\n")
	}

}
