package main

import (
	"fmt"
	"math"
)

// 浮点型
// Go语言支持两种浮点型数：float32、float64
// float32的浮点数最大范围约为：3.4e38，可以使用常量定义：math.MathFloat32
// float64的浮点数最大范围约为：1.8e308，可以使用常量定义：math.MathFloat64

func main() {
	f1 := math.MaxFloat32
	f2 := math.MaxFloat64
	f3 := 1.23456
	fmt.Printf("f1的类型是: %T\n", f1)
	fmt.Printf("f2的类型是: %T\n", f2)
	fmt.Printf("f3的类型是: %T\n", f3)
	f4 := float32(1.2345)
	fmt.Printf("f4的类型是: %T\n", f4)
}
