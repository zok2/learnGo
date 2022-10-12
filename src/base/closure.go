package main
// go 闭包实现  闭包 = 函数+引用环境
import "fmt"

func add() func(y int) int  {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main()  {
	f := add()
	sum := f(10)
	fmt.Println(sum)
	sum = f(20)
	fmt.Println(sum)
}