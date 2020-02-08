package main

import "fmt"

// 切片（slice）
// 切片是一个拥有相同类型元素的可变长度序列。它是基于数组类型做的一层封装
// 切片是一个引用类型，它的内部结构包括地址、长度、容量。
// 切片一般用于快速的操作一块数据集合
func main() {
	// 切片的定义
	var s1 []int // 定义一个存放int类型的切片（只限制类型，不限制长度）
	var s2 []string
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "天津", "上海"}
	fmt.Println(s1, s2)
	fmt.Printf("%T\t%T\n", s1, s2)
	// 切片的长度和容量
	fmt.Println("s1的长度", len(s1), "s2的长度", len(s2))
	fmt.Println("s1的容量", cap(s1), "s2的容量", cap(s2))
	// 由数组得切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s3 := a1[0:5] // [1 2 3 4 5]：基于一个数组切割，左闭右开
	s4 := a1[:5]  // a1[0:5]
	s5 := a1[3:]  // a1[0:len(a1)]
	s6 := a1[:]   // a1[0:len(a1)]
	fmt.Println(s3, s4, s5, s6)
	// 切片长度是指切片的第一个元素对应底层数组的元素到底层数组的最后一个元素的长度
	fmt.Println("s3的长度", len(s3), "s4的长度", len(s4), "s5的长度", len(s5), "s6的长度", len(s6))
	fmt.Println("s3的容量", cap(s3), "s4的容量", cap(s4), "s5的容量", cap(s5), "s6的容量", cap(s6))
	// 切片再切片
	s7 := s5[3:]
	fmt.Println("s7的长度", len(s7), "s7的容量", cap(s7))
	// 切片是一个引用类型，都指向了底层的数组
	a1[7] = 1300
	fmt.Println(s5, s7)
}
