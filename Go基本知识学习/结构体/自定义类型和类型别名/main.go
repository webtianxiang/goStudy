package main

import "fmt"

// 自定义类型和类型别名
// type后面跟的是类型
type myint int

// 类型别名：只在编译时有效，编译结束就失效了 
type yourint = int

func main() {
	var n myint
	var m yourint
	n = 100
	m = 100
	fmt.Printf("自定义类型：%T,值：%v\n", n, n)
	fmt.Printf("类型别名：%T,值：%v\n", m, m)
	var c rune
	c = '中'
	fmt.Printf("类型别名：%T,值：%v\n", c, c)
}
