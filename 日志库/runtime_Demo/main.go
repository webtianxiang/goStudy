package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	f2()
}
func f2() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
	fmt.Println(funcName)
}
func main() {
	f1()
}
