package main

import "fmt"

// map是一种无序的基于key-value的数据结构
// Go语言中，map是引用类型，需要初始化才可以使用
func main() {
	// 使用语法：map[keyType]valueType
	var m1 map[string]int
	fmt.Println(m1, m1 == nil)    // 没有初始化（还没有在内存中开辟空间）
	m1 = make(map[string]int, 10) // 要估算好map容量，避免在程序运行期间动态扩容
	m1["test1"] = 18
	m1["test2"] = 20
	m1["test3"] = 30
	fmt.Printf("v:%v 类型:%T\n", m1, m1)
	value, isOk := m1["test4"] // 如果不存在这个key，则拿到这个类型的零值
	if !isOk {
		fmt.Println("查无此值", value)
	}
	// map的遍历
	// 只遍历key
	for m2 := range m1 {
		fmt.Printf("key:%v val:%d t:%T\t", m2, m1[m2], m2)
	}
	fmt.Println()
	// 遍历key value
	for m2k, m2v := range m1 {
		fmt.Printf("key:%v val:%d t:%T\t", m2k, m2v, m2v)
	}
	fmt.Println()
	// 只遍历value
	for _, m2v := range m1 {
		fmt.Printf("val:%d t:%T\t", m2v, m2v)
	}
	fmt.Println()
	// map的删除
	delete(m1, "test1")
	fmt.Println(m1)
	delete(m1, "test4") // 若map无此元素或为nil，则不作操作
}
