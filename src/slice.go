package main

import "fmt"

/**
 *slice  学习
 */
func main() {
	city := []string{"北京", "上海", "广州"}
	city = append(city,"深圳")

	city = append(city[1:],city[3:]...)
	fmt.Println(city)
	fmt.Printf("city : %T\n", city)
	fmt.Printf("city len: %d\n", len(city))
	fmt.Printf("city cap: %d\n",cap(city))
	fmt.Printf("city addr: %p\n",city)
	code := make([]int,3,100)
	fmt.Printf("code : %T\n", code)
	fmt.Printf("code len: %d\n", len(code))
	fmt.Printf("code cap: %d\n",cap(code))
	fmt.Printf("code addr: %p\n",code)


	/*
	切片基于数组 ，切片引用类型
	初始化定义并赋值city
	地址三要素：地址，大小，容量
	切片是引用类型 获取切片的指针 %p + 变量名
	append 追加元素
	len 获取切片长度
	cap 获取切片容量 切片自动扩容
	获取切片的三种方式：
	直接声明复制
	基于数组得到切片
	基于切片得到切片
	[:] : 左开始的位置，:数量 左包含右不包含
	make ：引用类型做初始化 类型，len，cap
	new ：用来创建值类型
	copy 复制切片
	切片的变量 for   for range
	*/

}
