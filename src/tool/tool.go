package tool

import "fmt"

func Min(x ,y int) int {
	if x<y {
		return x
	}
	return y
}
func Max(x ,y int) int {
	if x>y {
		return x
	}
	return y
}
func Abs(x int) int {
	if x < 0  {
		return x * -1
	}
	return x
}
func init()  {
	fmt.Println("初始化。。。")
}