package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统可以查看、增加、删除学生
*/
var (
	students map[int64]*student
)

type student struct {
	id   int64
	name string
}

// 构造学生的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudents() {
	if len(students) == 0 {
		fmt.Println("温馨提示，暂无学生！")
	} else {
		for k, v := range students {
			fmt.Printf("学号: %v, 姓名: %s\n", k, v.name)
		}
	}
	fmt.Println("————————————————————————————————————————————————————")
}
func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanln(&name)
	newStu := newStudent(id, name)
	students[id] = newStu
	fmt.Println("添加成功！！！")
	fmt.Println("————————————————————————————————————————————————————")
}
func deleteStuden() {
	var (
		deleteID int64
	)
	fmt.Print("请输入要删除的学号：")
	fmt.Scanln(&deleteID)
	if _, v := students[deleteID]; !v {
		fmt.Println("******温馨提示******您输入的学好不存在，请重新输入：")
		deleteStuden()
	} else {
		delete(students, deleteID)
	}
}
func main() {
	// 开辟内存空间
	students = make(map[int64]*student, 100) // 变量声明
	// 1.打印菜单
	for {
		fmt.Println("欢迎使用学生管理系统")
		fmt.Println(`
				1.查看所有学生
				2.新增学生
				3.删除学生
				4.推出`)
		fmt.Print("请输入你要干什么？")
		// 2.等待用户选择下一步要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你选择了", choice)
		// 3.执行对应的函数
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudent()
		case 3:
			deleteStuden()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("请重新选择:")
		}
	}
}
