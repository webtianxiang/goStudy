package main

import "fmt"

// 关系运算符（==、！=、<=、>=、<、>）

func main() {
	var (
		a = 10
		b = 2
	)
	fmt.Println(a == b) // Go是强类型语言，相同类型的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
}
