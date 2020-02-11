package main

import "fmt"

// fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分
// fmt标准库提供的几个输出方法：print（给啥输出啥）、println（给啥输出啥+换行输出）、printf（格式化字符串输出）
//
func main() {
	// 获取输入
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容是：", s)
	var (
		name string
		age  int
		sex  string
	)
	// fmt.Scanf("%s %d%s\n", &name, &age, &sex)
	fmt.Scanln(&name, &age, &sex)
	fmt.Println(name, age, sex)
}

func output() {
	// 输出
	var n = 100
	fmt.Printf("%T\n", n) // 查看类型
	fmt.Printf("%v\n", n) // 查看值
	fmt.Printf("%d\n", n) // 十进制数
	fmt.Printf("%b\n", n) // 二进制数
	fmt.Printf("%o\n", n) // 八进制数
	fmt.Printf("%x\n", n) // 16进制数
	var name = "tianxiang"
	fmt.Printf("%s\n", name) // 字符串
	fmt.Printf("%v\n", name)
	fmt.Printf("%#v\n", name)
	// 整数->字符或者字符串
	fmt.Printf("%q\n", n)
	// 科学计数法
	fmt.Printf("%e\t%E\n", 110e129, 110E129)
	// %+v：按照值原样输出，但是结构体会额外打印结构体类型
	// %#：值结果更详细
	// %%：就是百分号
}
