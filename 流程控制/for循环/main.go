package main

import "fmt"

// for循环
func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	// 变种一
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("%d ", j)
	}
	fmt.Println()
	// 变种二
	var k = 5
	for k < 10 {
		fmt.Printf("%d ", k)
		k++
	}
	fmt.Println()
	// 无限循环
	// for {
	// 	fmt.Println()
	// }
	// for循环可以通过break、goto、return、panic语句强制跳出循环

	// for range
	// Go语言中可以使用for range来遍历数组、切片、字符串、map及通道
	// 通过for range遍历有以下特性：
	// 1.数组、字符串、切片返回索引和值
	// 2.map返回键和值
	// 3.通道（channel）只返回通道内的值
	s := "tian田翔xiang"
	for j, c := range s {
		fmt.Printf("%v %c\n", j, c)
	}
}
