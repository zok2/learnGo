package main

import (
	"fmt"
	"math"
)

func reverse(x int) (rev int)  {
	for x !=0 {
		if rev > math.MaxInt32/10 || rev < math.MinInt32 / 10 {
			return 0
		}
		digit := x %10
		x = x/10
		rev = rev*10+digit
	}
	return
}
func main()  {
	x := 123
	rev := reverse(x)
	fmt.Println(rev)
}