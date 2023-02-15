package main

import "fmt"

func main() {
	// 固定长度的数组，默认情况下，数组的每个元素都被初始化为元素类型对应的零值 int 0 ,string '' bool false 等
	var cityCode = [3]int{}
	var city = [3]string{"北京", "上海", "广州"}
	var citys = [3][3]string{{"北京", "上海", "广州"}, {"北京", "上海", "广州"}, {"北京", "上海", "广州"}}

	//...”省略号数组根据初始化述职来计算
	//初始化索引的顺序是无关紧要的，而且没用到的索引可以省略，和前面提到的规则一样，未指定初始值的元素将用零值初始化
	r := [...]int{10: -1}

	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println("cityCode :", cityCode)

	fmt.Println("city :", city)
	fmt.Println("city len :", len(city))
	fmt.Printf("city type = %T\n", city)

	fmt.Println("citys :", citys)
	fmt.Println("citys len :", len(citys))
	fmt.Printf("citys type = %T\n", citys)

	fmt.Println("r", r)
	fmt.Printf("r type = %T\n", r)

	fmt.Println("symbol", symbol) // "3 ￥"

	//一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，这时候我们可以直接通过==比较运算符来比较两个数组，
	//只有当两个数组的所有元素都是相等的时候数组才是相等的。不相等比较运算符!=遵循同样的规则。
	if city == citys[0] {
		fmt.Println("city  == citys ", true)
	} else {
		fmt.Println("city  != citys ", false)
	}
	//q := [3]int{1, 2, 3}
	//q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
	//p := [4]int{1, 2, 3, 4}
	// fmt.Println(q == p) //invalid operation: q == p (mismatched types [3]int and [4]int)
	//数组没有删除操作
}
