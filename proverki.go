package main

import "fmt"

func main() {
	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string
	fmt.Println(`
	pointerY
	pointerX
	pointerZ
	`)

	v := 12
	b := 13
	c := 14
	fmt.Println()
}
