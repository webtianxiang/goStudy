package main

import "fmt"

// 使用make()函数来创建切片
// 第一个参数是切片的类型、第二个参数时切片的长度、第三个参数是切片的容量
func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1:%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2:%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))
	// 切片的遍历
	// 索引遍历
	for i := 0; i < len(s1); i++ {
		fmt.Print(s1[i])
	}
	// for range循环
	for i, v := range s1 {
		fmt.Println("i:", i, "v:", v)
	}
}
