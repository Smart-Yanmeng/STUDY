package main

import "fmt"

// Teacher 结构体，将老师的各个属性统一放入结构体
type Teacher struct {
	// 变量名大写，让外界可以访问这个属性
	Name   string
	Age    int
	School string
}

func main() {
	var t1 Teacher
	fmt.Println(t1)

	t1.Name = "York"
	t1.Age = 18
	t1.School = "qztc"
	fmt.Println(t1)
}
