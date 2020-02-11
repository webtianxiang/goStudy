package main

import "fmt"

// 函数类型
func main() {
	a := f1
	fmt.Printf("v:%v T:%T\n", a, a)
	b := f2
	fmt.Printf("v:%v T:%T\n", b, b)
	f3(f2)
	fmt.Printf("v:%v T:%T\n", f4, f4)
	fmt.Println(f5(f2))
}
func f1() {
	fmt.Println("测试")
}
func f2() int {
	return 2
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	res := x()
	fmt.Println(res)
}
func f4(x, y int) int {
	return x + y
}

// 函数还可以作为返回值
// 高阶函数
func f5(x func() int) func(int, int) int {
	res := x()
	return func(x int, y int) int {
		return res * 2
	}
}
