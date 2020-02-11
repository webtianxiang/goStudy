package main

import "fmt"

// append()为切片追加元素
// Go语言的内建函数append()可以为切片动态添加元素。
// 每个切片指向一个底层数组，这个数组可以容纳一定量的元素，
// 当底层数组不能容纳新的元素时，切片就会自动按照一定的策略进行“扩容”
// 此时切片指向的底层数组就会更换。
// ”扩容“常常发生在append()函数调用时
func main() {
	var s1 []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("s1:%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// 调用append()函数必须用原来的切片变量接收返回值
	s1 = append(s1, 6, 7) // append()追加元素时，当底层数组放不下时，Go底层会把底层数组换一个
	fmt.Println(s1)
	fmt.Printf("s1:%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := []int{1, 2, 3, 4}
	s1 = append(s1, s2...) // ...表示拆开
	fmt.Printf("s1:%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
