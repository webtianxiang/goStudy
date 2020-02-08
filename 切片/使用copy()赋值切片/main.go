package main

import "fmt"

// copy()
// 第一个参数为：目标切片；第二个参数为：数据源切片
func main() {
	a1 := []int{1, 2, 3}
	a2 := a1
	var a3 = make([]int, 3, 5)
	// 使用copy()函数将a1中的元素复制到a3
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)
}
