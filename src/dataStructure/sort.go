package stack

import "fmt"

func main() {

	number :=[...] int{1,0,0,0,8,0,10,-2,-1,-9}
	j := 0
	for  i := 0; i< len(number); i++ {
		if number[i]!=0 {
			number[j]=number[i]
			j++
		}
	}

	for j:=j;j<len(number);j++ {
		number[j]=0
	}
	fmt.Print(number)
}
