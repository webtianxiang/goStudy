package main

import "fmt"

// 结构体是值类型
type person struct {
	name, gender string
}

// Go语言中函数参数永远是副本（改变的是副本）
func f1(p person) person {
	p.gender = "女"
	return p
}

// &:取地址；*:取地址对应的值
func f2(p *person) person {
	(*p).gender = "女" // 根据地址找到对应的原变量，修改的就是原变量
	p.gender = "女"    // 语法糖，自动根据.前的变量去判断，若为指针自动取其值
	return *p
}
func main() {
	var p person
	p.name = "田翔"
	p.gender = "男"
	fmt.Println(p)
	fmt.Println(f1(p), p)
	fmt.Println(f2(&p), p)
	// 结构体指针1
	// 我们可以通过new创建一个结构体类型指针
	var p2 = new(person) // p2本身存的是一个地址，而&p2是取p2这个变量的内存地址
	fmt.Printf("P2 type:%T value:%v\n", p2, p2)
	(*p2).name = "王紫冰"
	p2.gender = "女" // 用了语法糖，Go语言底层帮你转换为上面的形式：(*p2).
	// 结构体（指针）2.1：key-value初始化
	// var p3 = person{
	// 	name:   "崔洵",
	// 	gender: "男",
	// }
	var p3 = &person{
		name:   "崔洵",
		gender: "男",
	}
	fmt.Printf("%#v %T\n", p3, p3)
	// 结构体（指针）2.2：使用值列表的形式初始化
	// p4 := person{
	// 	"牛志杰",
	// 	"男",
	// }
	p4 := &person{
		"牛志杰",
		"男",
	}
	fmt.Printf("%#v %T\n", p4, p4)
	// 结构体内部字段占一个连续的内存内置
	fmt.Printf("name:%#v gender:%#v", &p4.name, &p4.gender)
}
