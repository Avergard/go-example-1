package main

import "fmt"

func main() {
	x := [13]int{1, 2, 3, 4, 5, 6, 7, 12: 8, 8: 12, 7: 1, 9: 4}
	fmt.Println(x)
	y := [...]int{1, 2, 3, 4, 5, 6, 7, 12: 8, 8: 12, 7: 1, 9: 4}
	fmt.Println(y)
	z := [...]string{"lol", 12: "lol"}
	fmt.Println(z)
	h := [12][13]string{"lol"}
	fmt.Println(h)

}
