package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start of the program")
	defer fmt.Println("End of the program")
	makeCounter := func(x int) func() int {
		counter := x
		return func() int {
			counter++
			return counter

		}
	}
	counter := makeCounter(0)
	for j := 0; j < 10; j++ {
		result := counter()
		fmt.Println(result)
		if result > 5 {
			fmt.Println("break")
			break
		}
	}

}
