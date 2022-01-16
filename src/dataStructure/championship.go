package stack

import (
	"fmt"
	"math"
)

func main()  {
	arr :=[...] int{17,48,56,42,61,19,23,37,14,9,99,32,20,25,54,73}
	a := int(math.Log2(float64(len(arr))))
	p := 1
	for i := 0;i<a;i++ {
		q := 2*p
		for t := 0;t< len(arr);t = t+q {
			if arr[t] < arr[t+p] {
				fmt.Println(arr[t] , arr[t+p])
				for j := 0;j<i;j++ {
					tmp := arr[t+j]
					arr[t+j] = arr[t+j+p]
					arr[t+j+p] = tmp
				}

				tmp := arr[t+p]
				arr[t+p]=arr[t+i]
				arr[t+i] =tmp
				fmt.Println(arr)
			}
		}
		p=q
	}
}