package main

import "fmt"

// 接口是一种类型，是一种特殊的类型，它规定了变量有哪些方法
// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。
// 接口的实现，如果一个变量先了该接口定义的所有的方法，则可以称该变量为该接口类型的变量
type speaker interface {
	speak()
}

// 引出接口的实例
type cat struct{}
type dog struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}
func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

// 在编程中会遇到以下场景，不关心一个变量是什么类型，我只关心能调用它的什么方法
func hit(x speaker) {
	// 接收一个参数，餐传进来谁就打谁
	x.speak()
}
func main() {
	var c1 cat
	var d1 dog

	hit(c1)
	hit(d1)
}
