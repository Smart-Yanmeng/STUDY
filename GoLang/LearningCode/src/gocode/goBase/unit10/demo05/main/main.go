package main

import (
	"fmt"
	"gocode/goBase/unit10/demo05/model"
)

func main() {
	// 创建一个 person 结构体实例
	p := model.NewPerson("York", 18)
	p.SetAge(20)

	fmt.Println((*p).GetAge())
}
