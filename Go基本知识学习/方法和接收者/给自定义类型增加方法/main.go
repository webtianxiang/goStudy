package main

import "fmt"

// 给自定义类型添加方法
// 不能给别的包中的类型添加方法，只能给自己的包中的自定义类型增加法法
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100)
	m.hello()
}