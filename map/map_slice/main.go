package main

import "fmt"

func main() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 0, 10)
	s1 = append(s1, make(map[int]string, 10))
	s1[0][10] = "test"
	fmt.Println(s1)
	var s2 = make([]map[int]string, 10, 10)
	s2[0] = make(map[int]string, 2)
	s2[0][18] = "18岁"
	fmt.Println(s2)
	// 值为切片的map
	var s3 = []int{1, 2, 3}
	m1 := make(map[string][]int, 10)
	m1["test"] = s3
	fmt.Println(m1)
}
