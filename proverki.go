package main

import "fmt"

type q struct {
	number1 int
	number2 int
	number3 int
}

func main() {
	number := q{
		number1: 111,
		number2: 112,
		number3: 113,
	}

	var x = []int{12, 13, 14, 15, 16}
	var y = [5]int{1, 2, 3, 4, 5}
	z := map[int][]int{
		1: []int{101, 102, 103, 104, 105},
		2: []int{106, 107, 108, 109, 110},
	}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(q{
		number1: 111,
		number2: 112,
		number3: 113,
	})
	fmt.Println(number)
}
