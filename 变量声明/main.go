package main

import "fmt"

// 声明变量并指定变量类型（在变量的生命周期内，变量类型不可变）
var name string
var age int
var isOk bool

// 批量声明（默认赋值为0值）
// var (
// 	name1 string // ""
// 	age1  int    // 0
// 	isOk1 bool   // false
// )

func main() {
	name = "Tianxiang"
	age = 23
	isOk = true
	fmt.Printf("name: %s\n", name)
	fmt.Printf("age: %d\n", age)
	fmt.Println("isOk:", isOk)
}
