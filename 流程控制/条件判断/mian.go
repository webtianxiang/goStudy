package main

import "fmt"

// if条件判断

func main() {
	age := 23
	// 单个条件
	if age > 18 {
		fmt.Println("您的年龄大于18岁")
	} else {
		fmt.Println("您的年龄小于18岁")
	}
	// 多个条件判断
	if age > 35 {
		fmt.Println("您的年龄大于35岁")
	} else if age > 18 {
		fmt.Println("您的年龄大于18岁")
	} else {
		fmt.Println("您的年龄小于18岁")
	}
	ifDemo()
}

func ifDemo() {
	// age的作用域只在if条件内部使用
	if age := 23; age > 18 {
		fmt.Println(">18")
	} else {
		fmt.Println("<18")
	}
}
