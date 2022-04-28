package main

import (
	"../tool"
	//  匿名导入  _ "../tool"
	// mytool 别名  "../tool"
	// . 不建议使用 方法重名会出现问题  "../tool"
	"fmt"
)
func main()  {
	c := tool.Max(10,2)
	fmt.Println(c)
}