package main

import (
	"fmt"
)

func main() {
	//创建channel实际上就是在内存中实例化了一个hchan结构体，并返回一个chan指针
	//channle在函数间传递都是使用的这个指针，这就是为什么函数传递中无需使用channel的指针，而是直接用channel就行了，因为channel本身就是一个指针
	// 无缓冲通道
	ch1 := make(chan int)
	close(ch1)
	// 有缓冲通道
	ch2 := make(chan int, 10)
	//通过汇编分析，我们知道，最终创建 chan 的函数是 makechan
	//新建一个 chan 后，内存在堆上分配.
	go func() {
		ch2 <- 2 // a send statement
		//ch1 <- 1 // a send statement
		close(ch2)
	}()

	go func() {
		//	y := <-ch1 // a receive statement; result is discarded

		//	fmt.Println(y)
	}()
	x, _ := <-ch2 // a receive expression in an assignment statement

	fmt.Println(x)
}
