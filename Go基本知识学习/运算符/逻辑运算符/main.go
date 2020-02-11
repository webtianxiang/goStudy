package main

import "fmt"

// 逻辑运算符（&&、||、!）
// &&：逻辑AND运算符
// ||：逻辑OR运算符
// !：逻辑NOT运算符
func main() {
	var age = 23
	// 如果年龄大于等于18岁并且年龄小于40岁打印青年
	if age >= 18 && age < 40 {
		fmt.Println("年轻人")
	}
	if age < 18 || age > 40 {
		fmt.Println("弱势群体")
	}
	// not 取反
	fmt.Println(!(age > 18))
}
