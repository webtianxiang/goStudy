package main

// 匿名导入包：表示导入了包，没有用到包中的方法。只是用到该包中的init方法
// import _ "包路径"
// init方法：当所在包被引用时，程序运行期间会自动执行init函数，不可被调用，只能自动调用。没有参数，没有返回值。
// init函数执行时机：全局声明 --> init() --> main()
// 如果多个包都定义了init()方法，最后被引入的包的init()先执行，以此内推，直至执行完，再执行main函数
import (
	// 引用目录从goPath的src目录下开始找
	// 起别名
	"goStudy/Go基本知识学习/包package/test"
	test2 "goStudy/Go基本知识学习/包package/test"
	// go语言中禁止循环导入
)

// 定义包名
// package test
func main() {
	test.Test()
	test2.Test()
}
