package main

// 1.make和new都是用来申请内存的
// 2.new很少用，一般用来给基本数据类型申请内存，返回也是该基本数据类型的指针类型
// 3.make也是用于分配内存的，区别于new，它只用于slice、map和channel的内存创建，
// 而且他的返回类型就是这三个基本类型本身，而不是指针类型，因为这三个类型本身就是引用类型
func main() {
	// 使用方法
	// make(t Type, size...IntegerType)Type
}
