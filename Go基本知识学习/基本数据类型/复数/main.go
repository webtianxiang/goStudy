package main

import "fmt"

// 复数可以分为complex64、complex128
// 复数有实部和虚部，complex64的实部和虚部均为32位；complex128的实部和虚部均为64位

func main() {
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 4i
	fmt.Println(c1, c2)
}
