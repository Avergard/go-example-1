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
	counter := makeCounter(0)
	for i := 0; i < 10; i++ {
		result := counter()
		fmt.Println(result)
		if result > 5 {
			fmt.Println("Break")
			break
		}
	}
}
