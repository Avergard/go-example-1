package main

import "fmt"

func main() {
	defer fmt.Println("End of the program")
	fmt.Println("Start of the program")
	makeCounter := func(x int) func() int {
		counter := x
		return func() int {
			counter++
			return counter
		}
	}
}
