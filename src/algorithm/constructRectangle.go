package main

import (
	"fmt"
)

func spiralOrder(matrix [][]string,n int )  {
	//var l,w []int
	//for k,v := range matrix{
	//	for index,val := range v{
	//		if val == "*" && len(l) {
	//			x[0] = index
	//		}
	//
	//	}
	//}
}

func main()  {
    mat := [][] string {
    	{"-","-","*","-"},
		{"-","-","-","-"},
		{"*","-","-","-"},
		{"-","-","-","-"},
    }
	var w,l []int
	
	fmt.Println(l,len(l))
	fmt.Println(w,len(w))
	var n = 4
	spiralOrder(mat,n)

}