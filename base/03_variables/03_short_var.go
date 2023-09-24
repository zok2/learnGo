package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%T | %v \n", a, a)
	fmt.Printf("%T | %v \n", b, b)
	fmt.Printf("%T | %v \n", c, c)
	fmt.Printf("%T | %v \n", d, d)
	fmt.Printf("%T | %v \n", e, e)
	fmt.Printf("%T | %v \n", f, f)
	fmt.Printf("%T | %v \n", g, g)

}
