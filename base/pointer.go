package main

import "fmt"

// 变量交换
func swap1 (a, b int) {
	var tmp int
	tmp = a
	a = b
	b = tmp
}
// 指针方式交换
func swap2(pa, pb * int)  {
	var tmp = *pa
	*pa = *pb
	*pb = tmp
}
func main()  {
	var a,b int = 10,20
	swap2(&a,&b)
	fmt.Println(&a)
	// pp **int 二级指针
}