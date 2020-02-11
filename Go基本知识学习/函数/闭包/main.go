package main

import "fmt"

// 闭包
// 闭包是一个函数，这个函数包含了外部作用域的变量

// 底层原理
// 1.函数可以作为返回值
// 2.函数内部查找变量的顺序，现在内部找，找不到再向外部找
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
func f3(f func(int, int), x, y int) func() {
	temp := func() {
		f(x, y)
	}
	return temp
}

// 定义一个函数对f2进行包装
func main() {
	f3(f2, 2, 4)()
	f1(f3(f2, 2, 4))
	// f1(f2)
}
