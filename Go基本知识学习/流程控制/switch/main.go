package main

import "fmt"

// switch
// 简化大量的判断（一个变量和具体的值做比较）

func main() {
	// 基础用法
	var n = 2
	switch n {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println(0)
	}
	fmt.Println()
	// 变种二
	switch m := 2; m {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println(0)
	}
	fmt.Println()
	// 变种三
	switch n {
	case 1, 2, 3, 4:
		fmt.Println(n)
	case 5, 6, 7, 8:
		fmt.Println(n)
	default:
		fmt.Println(0)
	}
	fmt.Println()
	// 变种四
	switch {
	case n > 1 && n < 5:
		fmt.Println(n)
	case n > 5:
		fmt.Println(n)
	default:
		fmt.Println(0)
	}
	fmt.Println()
}
