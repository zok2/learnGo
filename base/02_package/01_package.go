package main

import (
	"fmt"
	"work/tool"
)

func main() {
	// 调用tool package
	var a, b = 1, 2
	min := tool.Min(a, b)
	fmt.Printf("min a,b  = %d", min)
}
