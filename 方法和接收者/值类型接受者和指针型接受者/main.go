package main

import "fmt"

type dog struct {
	name string
}
type person struct {
	name string
	age  int8
}

// 构造函数
func newDog(name string) *dog {
	return &dog{name: name}
}
func newPerson(name string, age int8) person {
	return person{name: name, age: age}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法具体的类型，它多用类型名首字母小写（形参）
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}

// 结构体是值类型，使用值接收者是传拷贝进入
func (p person) guonian() {
	p.age++
}

// 使用指针接收者，传内存地址进去
func (p *person) guoshengri() {
	// p.age++
	(*p).age++
}
func (p *person) dream() {
	fmt.Println("不上班也能挣钱")
}
func main() {
	d1 := newDog("旺财")
	d1.wang()
	p1 := newPerson("王紫冰", 23)
	fmt.Println(p1)
	p1.guonian()
	fmt.Println(p1)
	p1.guoshengri()
	fmt.Println(p1)
	p1.dream()
}

// 什么时候使用指针型接收者
// 1.当我需要修改接收者中的值
// 2.接收者是拷贝成本大，是大对象时
// 3.保证一致性时（当代码中有一个使用指针型接收者时，后面的其他方法也尽量保持统一使用指针型接收者）
