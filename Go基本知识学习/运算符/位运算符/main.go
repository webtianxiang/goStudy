package main

import "fmt"

// 位运算符（针对二进制数）
// &：按位与(两位均为一才为一)
// |：按位或(两位有其一则为一)
// ^：按位异或(两位不一样为一)
// <<：将二进制左移指定位数
// >>：将二进制右移指定位数

func main() {
	// 5的二进制：101
	// 2的二进制：10
	fmt.Println(5 & 2)
	fmt.Println(5 | 2)
	fmt.Println(5 ^ 2)
	fmt.Println(5 << 2)
	fmt.Println(5 >> 2)
}
