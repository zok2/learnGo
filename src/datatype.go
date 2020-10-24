package main

import (
	"fmt"
	"math"
	"strings"
)

func main()  {
	var a int =10
	var b int = 077
	fmt.Println(a,b)
	fmt.Printf("%d\n",a)
	fmt.Printf("%b\n",a)
	fmt.Printf("%o \n",a)
	fmt.Printf("%p \n",&a)

	fmt.Printf("%f \n",math.MaxFloat32)
	fmt.Printf("%f \n",math.MaxFloat64)

	fmt.Println("\"/home/img\"")

	var name = "星光"
	var detailed = `asd
asd
`
	fmt.Println(name)
	fmt.Println(detailed)
	fmt.Println(len(name))
	fmt.Println(len(detailed))
	fmt.Println(name+"111"+detailed)
	fmt.Println(fmt.Sprintf("%s---+%s",name,detailed))
	s1 := strings.Split(detailed,"a")
	fmt.Println(strings.Split("a,b,c",","))
	fmt.Println(s1)

	//是否包含 返回true false
	fmt.Println(strings.Contains(detailed,"a"))
	fmt.Println(strings.HasPrefix(detailed,"a"))
	fmt.Println(strings.HasSuffix(detailed,"\n"))

	fmt.Println(strings.Index(detailed,"s"))
	fmt.Println(strings.LastIndex(detailed,"s"))

	fmt.Println(strings.Join(strings.Split("a,b,c",","),"-"))
}