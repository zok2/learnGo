package main

import (
	"fmt"
	"math"
)
// 通项公式
func climbStairs1(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
	pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
	return int(math.Round((pow1 - pow2) / sqrt5))
}


func climbStairs2(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main()  {
	res1 := climbStairs1(100)
	res2 := climbStairs2(100)
	fmt.Println(res1)
	fmt.Println(res2)
}