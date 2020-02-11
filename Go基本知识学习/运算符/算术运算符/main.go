package main

import "fmt"

// 算数运算符（+、-、*、/、%）

func main() {
	var (
		a = 10
		b = 6
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // 单独的语句不能放到右边赋值  => a = a + 1
	b--
	fmt.Println(a, b)
}
