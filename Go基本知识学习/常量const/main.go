package main

import "fmt"

// 常量
// 程序运行期间不会改变的量
const pi = 3.1415926

// 批量声明常量
const (
	STATUS_OK = 200
	NOT_FOUND = 404
)

// n2，n3默认跟随n1赋值
const (
	n1 = 100
	n2
	n3
)

// iota常量计数器
// iota在const关键字出现时，会讲iota计数初始为0，const每新增一行常量声明iota计数一次
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)
const (
	b1 = iota
	b2
	_
	b3
)
const (
	c1 = iota
	c2 = 100
	c3
	c4 = iota
	c5
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n:", n1, n2, n3)
	fmt.Println("a:", a1, a2, a3)
	fmt.Println("b:", b1, b2, b3)
	fmt.Println("c:", c1, c2, c3, c4, c5)
	fmt.Println("d:", d1, d2, d3, d4)
	fmt.Println("数量级：", KB, MB, GB, TB, PB)
}
