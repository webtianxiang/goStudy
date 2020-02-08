package main

import "fmt"

// 数组：数组是同一种类型元素的集合
// 在Go语言中，从数组的声明开始要指定数组长度 ，就可以在使用过程中修改数组成员，但是数组大小不可变化
// 数组的长度是数组类型的一部分
func main() {
	//  数组声明、定义
	var a1 [3]bool
	var a2 [4]bool
	fmt.Println(a1, a2)
	fmt.Printf("%T %T", a1, a2)

	// 数组初始化
	// 如果不初始化，默认元素都为“零值”（布尔值：false；整型和浮点型：0；字符串：""）
	// 初始化方式一
	var a3 = [3]bool{false, true, false}
	fmt.Println(a3)
	// 初始化方式二
	// var a4 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a4 = [...]int{0, 1, 2, 3, 4, 5, 6} // 根据初始值自动推断数组长度
	fmt.Println(a4)
	// 初始化方式三
	var a5 = [5]int{1, 2}
	var a6 = [5]int{0: 1, 4: 2}
	fmt.Println(a5, a6)
}
