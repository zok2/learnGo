package main

import "fmt"

func main()  {
	// 固定长度的数组，默认值 0，值类型
	var city [3]string= [3]string{"北京","上海","广州"}
	fmt.Println(city)
	fmt.Println(len(city))
	//长度不同类型不同
	fmt.Printf("city type = %T\n",city)


	r:=[...] int{10:-1}
	fmt.Printf("r type = %T\n",r)
	fmt.Println(r)
	//1.求出数组中的和
	a :=[...] int{1,3,5,7,8,0,10,-2,-1,-9}
	sum:=0
	for _,v := range a {
		sum+=v
	}
	fmt.Println(sum)
	//2. 找出和等8的下标
	other:=8
	for key,value := range a {
		for k,v := range a{
			if v+value== other{
				println(key,k)
			}	
		}
	}
}
