package main

import "fmt"

func main() {

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
		slice can only be compared to nil 切片只可以和nil比较
	*/
	//声明切片 ,没有分配内存 直接操作会造成 painc
	var citys []string
	// citys[0] = "北京"  // panic: runtime error: index out of range [0] with length 0

	//赋值,初始化
	citys = []string{"上海", "北京"}
	fmt.Println("声明切片citys ：", citys, "长度 len :", len(citys), "容量 cap：", cap(citys), citys == nil)

	//使用数组初始化
	cityCodes := [...]int{1, 2, 3, 4, 5, 6}
	codes := cityCodes[:]
	nums1 := cityCodes[0:3]
	nums2 := cityCodes[3:6]
	fmt.Println("声明切片codes ：", codes, "长度 len :", len(codes), "容量 cap：", cap(codes), codes == nil)

	// 闭区间 到开区间  [a:b] a<= a   b<b
	fmt.Println("声明切片nums1 ：", nums1, "长度 len :", len(nums1), "容量 cap：", cap(nums1), nums1 == nil)
	fmt.Println("声明切片nums2 ：", nums2, "长度 len :", len(nums2), "容量 cap：", cap(nums2), nums2 == nil)

	// 声明切片并初始化默认值（元素类型的0值）
	age := make([]int, 3, 3)
	fmt.Println("声明切片age ：", age, "长度 len :", len(age), "容量 cap：", cap(age), age == nil)

	//append  切片增加与删除
	codes = append(codes, 7)
	fmt.Println("切片codes ：", codes, "长度 len :", len(codes), "容量 cap：", cap(codes), codes == nil)
	codes = append(codes, nums1...)
	fmt.Println("切片codes ：", codes, "长度 len :", len(codes), "容量 cap：", cap(codes), codes == nil)
	codes = append(codes[:4], codes[6:]...)
	fmt.Println("切片codes ：", codes, "长度 len :", len(codes), "容量 cap：", cap(codes), codes == nil)

	//copy
	//浅拷贝 赋值的是地址
	nums3 := nums1
	fmt.Printf("nums1 地址 %p\n", nums1)
	fmt.Printf("nums3 地址 %p\n", nums3)

	//var nums4 []int  copy时需要创建内存空间
	var nums4 = make([]int, 8)
	copy(nums4, nums1)
	fmt.Printf("nums4 地址 %p\n", nums4)
	fmt.Println(nums4)
}
