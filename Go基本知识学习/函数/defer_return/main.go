package main

import "fmt"

// Go语言中的return不是原子操作，在底层分两步执行
// 第一步：将返回值赋值
// 第二步：真正的res返回

// 函数中如果存在defer，那么defer定义的语句，执行的时机实在第一步和第二步之间
func main() {
	fmt.Println("f1", f1())
	fmt.Println("f2", f2())
	fmt.Println("f3", f3())
	fmt.Println("f4", f4())
}
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++ // 函数传参拷贝的副本，改变的是副本
	}(x)
	return 5
}
