package main

import "fmt"

// Go语言中不存在指针操作，只需要记住以下两个符号：
// 1.&：取地址
// 2.*：根据地址取值
func main() {
	n := 12
	fmt.Println(&n)
	p := &n
	fmt.Printf("值：%v 类型：%T\n", p, p) // *int：表示int类型的指针
	m := *p
	fmt.Println(m)
}
