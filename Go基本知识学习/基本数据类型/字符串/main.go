package main

import (
	"fmt"
	"strings"
)

// Go语言中的字符串以原生数据类型出现，使用字符串就像是使用其他uansheng数据类型（int、float32、float64、bool）一样。
// Go语言里的字符串的内部实现使用UTF-8编码。
// 字符串的值为双引号（""）中的内容，可以再Go计的源码中直接添加非ASCII码
// Go语言中单引号包裹的是字符

// 字符串转义符
// \r --> 回车符（返回行首）
// \n --> 换行符
// \t --> 制表符
// \' --> 单引号
// \" --> 双引号
// \\ --> 反斜杠
func main() {
	// 字符串
	s1 := "tianxiang"
	s2 := "田翔"
	// 字符（单独的字母、数字、汉字、符号标识字符）
	c1 := 'h'
	c2 := '1'
	c3 := '田'
	fmt.Printf("%s%s\n", s1, s2)
	fmt.Print(c1, c2, c3)
	// 声明多行字符串
	s3 := `
第一行
第二行
第三行`
	fmt.Printf("%s\n", s3)
	// 字符串相关操作
	// 字符串求长度
	fmt.Println(len(s2))
	fmt.Println(len("123"))
	// 字符串拼接
	fmt.Println(s1 + s2)
	s1s2 := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(s1s2)
	// 字符串分割
	res := strings.Split(s1, "i")
	fmt.Println(res)
	// 字符串包含
	fmt.Println(strings.Contains(s1, "tian"))
	// 前缀后缀判断
	fmt.Println(strings.HasPrefix(s1, "tian"))
	fmt.Println(strings.HasSuffix(s1, "tian"))
	// 获取字符串字符出现位置（下标）
	fmt.Println(strings.Index(s1, "i"))
	fmt.Println(strings.LastIndex(s1, "i"))
	// 数组拼接字符串
	fmt.Println(strings.Join(res, "+"))
}
