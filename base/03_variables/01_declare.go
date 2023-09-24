package main

import "fmt"

/*
var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明
*/
const boilingF = 212.0

// 输出水的沸点
func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("沸腾点 = %g°F Or %g°C\n", f, c)

	const freezingF = 32.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
