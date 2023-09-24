package main

import "fmt"

func main() {

	//var 变量名字 类型 = 表达式  “类型”或“= 表达式”两个部分可以省略其中的一个
	var name string = "zok"
	//var 变量名字  = 表达式 (通过初始化表达式来推导变量的类型信息)
	var age = 18
	// 省略 表达式  数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零值是空字符串，
	//接口或引用类型（包括slice、指针、map、chan和函数）变量对应的零值是nil。 （后面单独演示）
	var sex bool

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(sex)

}
