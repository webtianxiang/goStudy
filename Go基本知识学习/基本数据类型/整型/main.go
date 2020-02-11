package main

import "fmt"

// 整型
// 带符号int8、int16、int32、int64
// 无符号：uint8、uint16、uint32、uint64
// 特殊整型：int、uint（根据所在操作系统，32位即int32/uint32;64位即int64/64uint）、uintptr(无符号整型指针：用于存储一个指针)

// Go中无法直接声明二进制数，需要转化
func main() {
	var a int = 10
	fmt.Printf("a十进制：%d\n", a)  // 十进制打印
	fmt.Printf("a二进制：%b\n ", a) // 十进制转化为二进制打印
	fmt.Printf("a八进制：%o\n", a)  // 十进制转化为八进制打印
	fmt.Printf("a十六进制：%x\n", a) // 十进制转化为十六进制打印
	// 声明八进制
	b := 077
	fmt.Printf("b十进制：%d\n", b) // 十进制打印
	// 声明十六进制
	c := 0x123
	fmt.Printf("c十进制：%d\n", c)
	// 查看变量类型
	fmt.Printf("a的变量类型是：%T\n", a)
	fmt.Printf("b的变量类型是：%T\n", b)
	fmt.Printf("c的变量类型是：%T\n", c)
	// 明确指定变量类型，否则就是默认int类型
	d := int8(9)
	e := int16(15)
	fmt.Printf("d的变量类型是：%T\n", d)
	fmt.Printf("e的变量类型是：%T\n", e)
}
