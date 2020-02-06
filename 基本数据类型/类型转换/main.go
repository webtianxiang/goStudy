package main

import "fmt"

// Go语言中只有强制类型转换，没有隐式类型转换。
// 类型转换只有在两个类型支持类型相互转换的时候才可以使用。
// 布尔类型不支持类型转换
func main() {
	n := 10
	var f float64
	f = float64(n)
	fmt.Printf("%T", f)
	// 整型浮点型相互转换
	// 字符串型和rune类型相互转化
}
