package main

import (
	"fmt"
	"os"
	"strings"
)
// https://docs.hacknode.org/gopl-zh/ch1/ch1-01.html
func main()  {
	fmt.Println("hello  world! " + " balabala")
	fortest1()
	fortest2()
	fmt.Println(strings.Join(os.Args[1:], "+"));
	fmt.Println(os.Args[1:])
	test1()
	test2()
}

func fortest1()  {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
func fortest2()  {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep+arg
		sep = " "
	}
	fmt.Println(s)

}
func test1() {
	fmt.Println(os.Args[0:])
	fmt.Println(os.Args[0])
}
func test2()  {
	for k,v := range os.Args[1:]{
		fmt.Println(k,v,"|")
	}
}
