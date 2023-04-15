package model

import "fmt"

// 其他包不能直接访问
type person struct {
	name string
	age  int
}

// NewPerson - 定义工厂模式的函数，相当于构造器
func NewPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// SetAge - 定义 set 和 get 方法进行封装，确保被封装字段的安全合理性
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("对不起，你传入的年龄范围不正确")
	}
}
func (p *person) GetAge() int {
	return p.age
}
