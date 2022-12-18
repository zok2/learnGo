package main

import "fmt"

func foo1(a,b int) int {
	return a+b
}

//返回多个返回值
func foo2(a,b int)(int,int)  {
	return a,b
}
//返回多个返回值，有形参名。
func foo3(a,b int)(r1 ,r2 int)  {
	//r1 r2 属于返回形参， 默认初始化，默认0
	fmt.Println("r1 = ", r1,"r2 = ",r2)

	r1 = a
	r2 = b
	return r1,r2
}

func main()  {
	c := foo1(10,2)
	d,e := foo2(10,2)
	f,g := foo3(10,2)
	fmt.Println(c)
	fmt.Println(d,e)
	fmt.Println(f,g)
}



