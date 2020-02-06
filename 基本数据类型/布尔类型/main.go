package main

import "fmt"

// 布尔类型只有两个值：true、false
// 布尔类型变量的默认值为false
// Go语言中不允许将整数强制转化为布尔类型
// 布尔类型无法参与数值运算，也无法与其他类型进行转换

func main() {
	var b1 bool
	b2 := true
	fmt.Println(b1, b2)
}
