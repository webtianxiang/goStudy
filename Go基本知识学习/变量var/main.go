package main

import "fmt"

// 声明变量并指定变量类型（在变量的生命周期内，变量类型不可变）
// 同一作用域不能声明同名变量
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
	// 声明变零并赋值
	var s1 string = "hhh"
	fmt.Println(s1)
	// 类型推导，当未指定变量类型时，根据赋值隐式声明变量类型
	var s2 = "qqq"
	fmt.Println(s2)
	// 简短变量声明
	s3 := "www"
	fmt.Println(s3)
	// 匿名变量（匿名变量不占用命名空间，也不分配内存空间，所以不存在重复声明）
	x, _ := foo()
	_, y := foo()
	fmt.Println(x, y)
}

func foo() (int, string) {
	return 10, "bar"
}
