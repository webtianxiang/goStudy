package main

import "fmt"

// 函数

// 函数的定义
func sum(a int, b int) int {
	return a + b
}

// 没有返回值的函数
func f1(a int, b int) {
	fmt.Println(a, b)
}

// 没有参数但有返回值
func f2() int {
	return 3
}

// 返回值可以命名也可以不命名
// 命名的返回值就相当于在函数中声明了该变量
// 如果使用命名返回值，可以不写return后面的值
func f3() (res int) {
	res = 3
	return
}

// 多个返回值
func f4() (int, int) {
	return 1, 2
}

// 当参数中，连续两个参数类型一致时，我们可以将非最后一个参数类型省略
func f5(a, b int) int {
	return a + b
}

// 可变长参数， 必须放到函数参数的后面
func f6(a int, b ...int) {
	fmt.Println(a, b)
}

// 函数存在的意义？
// 1.函数是一段代码的封装。
// 2.把一段逻辑抽象出来封装到一个函数中，给他起一个名字，每次用到他的时候用函数名调用即可
// 3.使用函数能够让代码更清晰、更简洁
// Go语言中，没有函数默认参数的概念
func main() {
	fmt.Println(sum(2, 4))
	f1(2, 4)
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	f6(6, 1, 2, 3)
}
