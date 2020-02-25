package test

import "fmt"

// 包中的标识符（变量名、函数名、结构体、接口等），如果首字母是小写的，表示私有（只能在当前包使用）
// 首字母大写的标识符可以在其他中的包中使用

// Test 供其他包使用的方法
func Test() {
	fmt.Println("我是一个测试包")
}
func init() {
	fmt.Println("我是init函数，import导入我所在的包时，自动执行")
}
