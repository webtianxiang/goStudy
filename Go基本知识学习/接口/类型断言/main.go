package main

import "fmt"

// 类型断言1
func assign1(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串", str)
	}
}

// 类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch v := a.(type) {
	case string:
		fmt.Printf("传进来的是一个字符串：%v\n", v)
	case int:
		fmt.Printf("传进来的是一个int类型：%v\n", v)
	case bool:
		fmt.Printf("传进来的是一个bool：%v\n", v)
	default:
		fmt.Printf("传进来的是一个变量：%v\n", v)
	}
}
func main() {
	assign1(123)
	assign2(123)
}
