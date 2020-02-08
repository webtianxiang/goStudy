package main

import "fmt"

// 数组的遍历
func main() {
	var a1 [3]string = [3]string{"北京", "天津", "上海"}
	fmt.Println(a1)
	// 第一种遍历，根据索引遍历
	for i := 0; i < len(a1); i++ {
		fmt.Printf("%s ", a1[i])
	}
	// 第二种遍历，根据for range
	for i, s := range a1 {
		fmt.Printf("%d %s %s\n", i, s, a1[i])
	}
}
