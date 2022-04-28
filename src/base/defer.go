package main

import "fmt"

func end3()  {
	fmt.Println("main :: end3")
}
// defer 与 return  的先后顺序(先return 在defer)
func deferFunc()  {
	fmt.Println("defer Func")
}
func returnFunc() int {
	fmt.Println("return Func")
	return 0
}
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}
func main()  {
	// 类似析构函数，顺序问题，压栈的方式，结束前出栈
	defer end3()
	defer fmt.Println("main :: end2")
	defer fmt.Println("main :: end1")
	fmt.Println("main ::  hello go")
	
	returnAndDefer()
	
}
