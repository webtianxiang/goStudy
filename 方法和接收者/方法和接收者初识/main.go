package main

import "fmt"

// Go 语言中方法（method）是一种作用于特定类型变量的函数。这种特殊类型变量接受者。
// 接受者的概念就类似其他语言中的this或self
// 定义格式
// Go 语言中如果标识符首字母是大写，就表示对外部可见，是能被别的包调用的

// Dog 这是一个狗
type Dog struct {
	name string
}

type dog struct {
	name string
}

// 构造函数
func newDog(name string) *dog {
	return &dog{name: name}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法具体的类型，它多用类型名首字母小写（形参）
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}
func main() {
	d1 := newDog("旺财")
	d1.wang()
}
