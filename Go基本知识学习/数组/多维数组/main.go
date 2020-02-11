package main

import "fmt"

func main() {
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{0, 1},
		[2]int{2, 3},
		[2]int{4, 5},
	}
	fmt.Println(a11)
	// 多维数组遍历
	for i := 0; i < len(a11); i++ {
		fmt.Print(a11[i])
		for j := 0; j < len(a11[i]); j++ {
			fmt.Print(a11[i][j])
		}
		fmt.Println()
	}
	// 数组是值类型，赋值和传参会复制整个数组
	b1 := [3]int{0, 1, 2}
	b2 := b1
	b2[2] = 100
	fmt.Println(b1, b2)
}
