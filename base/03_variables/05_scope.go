package main

import (
	"fmt"
	"work/tool"
)

var g = "g"

// 作用域
func main() {
	f := "F"
	fmt.Println(f)
	fmt.Println(g)
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i] //x[i] 赋值到区块变量x
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println(x)
	case1()
	// 包作用域，方法名or 变量常量名首字母大写（大写公用、小写私有）
	fmt.Println(tool.InitStr)
	fmt.Println(tool.Max(1, 3))

}

// for 区块作用域
func case1() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
	fmt.Println(x)
}
