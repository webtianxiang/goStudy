package main

import "fmt"

func calc(index string, a, b int) int {
	res := a + b
	fmt.Println(index, a, b, res)
	return res
}
func testMain() {
	a := 1
	b := 2
	defer calc("1", a, calc("2", a, b))
	a = 0
	defer calc("3", a, calc("4", a, b))
	b = 1
}
