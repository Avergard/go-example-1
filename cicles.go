package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//if
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	f := rand.Intn(15)
	if f == 0 {
		fmt.Println(f, "zero")
	} else if f > 9 {
		fmt.Println(f, "mnogo")
	} else {
		fmt.Println(f, "malo")
	}

	//for
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
	//бесконечный цикл for
	/*
	for {
		fmt.Println("hello")

	}
*/
}
