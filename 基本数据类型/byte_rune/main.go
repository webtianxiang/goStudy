package main

import (
	"fmt"
)

// len()求得的是字符串转化byte类型长度
// Go语言中的字符有两种类型
// 一为uint8类型，或成为byte类型，代表了ASCII码的一个字符；二为rune类型，代表了一个UTF-8字符
// Go语言中为了定义非ASCII码字符定义了rune类型
// 当处理中文、日文、韩文等其他符合字符时，则需要用到rune类型。rune类型实际上是一个int32
// Go使用了特殊的rune类型来处理Unicode，让基于Unicode的文本处理更加方便。
func main() {
	s1 := "Tx田翔001"
	n := len(s1)
	i := 0
	for ; i < len(s1); i++ {
		fmt.Printf("%c ", s1[i])
	}
	fmt.Println()
	for _, c := range s1 {
		fmt.Printf("%c ", c)
	}
	fmt.Println()
	fmt.Printf("%d %d\n", n, i)

	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转化成了一个rune切片
	s3[0] = '红'
	fmt.Printf("%s %s\n", string(s3), s2) // 把rune切片强制转化为字符串

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T c2-10:%d\n", c1, c2, c2)

	t1 := "tian田翔"
	// t2 := "tian"
	// t3 := "田翔"
	t4 := byte('t')
	t5 := rune('田')
	fmt.Printf("t1:%d t2:%b t3:%b\n", len(t1), t4, t5)
	fmt.Printf("%T %T", t4, t5)
}
