package main

import "fmt"

// 对已经关闭的channel进行读写操作会发生什么?
func main() {
	fmt.Println("case1........")
	case1()
	fmt.Println("case2........")
	case2()
	fmt.Println("case3........")
	case3()
	fmt.Println("case4........")
	case4()
}

// 1. 读一个已经关闭的通道
func case1() {
	channel := make(chan int, 10)
	channel <- 2
	close(channel)
	x := <-channel
	fmt.Println(x)
}

// 遍历读关闭通道

func case2() {
	channel := make(chan int, 10)
	channel <- 2
	channel <- 3
	close(channel) //若不关闭通道，则会报死锁错误
	for num := range channel {
		fmt.Println(num)
	}
}

// 2. 写一个已经关闭的通道
func case3() {
	channel := make(chan int, 10)
	close(channel)
	//channel <- 1
}

/*[Output]: panic: send on closed channel*/

// 3. 关闭一个已经关闭的管道
func case4() {
	channel := make(chan int, 10)
	close(channel)
	//close(channel)
}

/*[Output]: panic: close of closed channel*/
