package main

import "fmt"

// var f1 = func(x, y int) int {
// 	return x + y
// }

func main() {
	f1 := func(x, y int) int {
		return x + y
	}
	fmt.Println(f1(2, 4))
	// 如果只是调用一次，可以写成立即调用函数
	func() {
		fmt.Println("我是匿名函数")
	}()
}
