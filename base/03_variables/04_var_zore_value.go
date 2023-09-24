package main

import "fmt"

func main() {

	// 省略 表达式  数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零值是空字符串，
	//接口或引用类型（包括slice、指针、map、chan和函数）变量对应的零值是nil。 （后面单独演示）
	var a string
	var b int
	var c bool
	var d float64

	fmt.Printf("%T | %v \n", a, a)
	fmt.Printf("%T | %v \n", b, b)
	fmt.Printf("%T | %v \n", c, c)
	fmt.Printf("%T | %v \n", d, d)
}
