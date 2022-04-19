package main

import "fmt"
//全局变量
var g = "全局变量"

/*
	四种变量声明方式
 */
func main()  {
	//声明一个变量，默认的值是0
	var a int
	fmt.Printf("type of a = %T\n",a)

	//声明一个变量，初始化一个值
	var b int = 1
	var c string = "abc"
	fmt.Printf("b = %d type of b = %T\n",b,b)
	fmt.Printf("c = %s type of c = %T\n",c,c)

	//声明一个变量，省去数据类型
	var d = 1
	var e = "abcd"
	fmt.Printf("d = %d type of d = %T\n",d,d)
	fmt.Printf("e = %s type of e = %T\n",e,e)

	//（常用的方法） 省去var 关键字，直接自动匹配
	f := 100
	fmt.Printf("f = %d type of f = %T\n",f,f)

	//全局变量 := 只能用在函数体内声明  var等方式用来声明全局变量
	fmt.Printf("g = %s type of g = %T\n",g,g)

	// 声明多个变量
	var x,y int = 100,200
	fmt.Println("x=" , x , "y" , y)
	var (
		i int = 100
		j bool = false
	)
	fmt.Println("i=" , i , "j" , j)

}
