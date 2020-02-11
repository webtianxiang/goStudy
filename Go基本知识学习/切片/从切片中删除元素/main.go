package main

import "fmt"

// Go语言中没有删除切片中的元素的方法，需要自己实现
func main() {
	// 删除元素3
	a1 := []int{1, 2, 3, 4, 5, 6}
	a1 = append(a1[:2], a1[3:]...)
	fmt.Println(a1)
	x1 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("x1:%v len(x1)=%d cap(x1)=%d\n", x1, len(x1), cap(x1))
	s1 := x1[:]
	s1 = append(s1[:3], s1[4:]...)
	fmt.Println(s1)
	fmt.Printf("x1:%v len(x1)=%d cap(x1)=%d\n", x1, len(x1), cap(x1))
}
