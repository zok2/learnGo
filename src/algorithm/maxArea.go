package main

import (
	"../tool"
	"fmt"
)
// 双指针计算最大面积
func maxArea(height []int) int {
	area,l,r :=0, 0, len(height)-1

	for l<r {
		tmp := tool.Min(height[l],height[r]) * (r-l)
		area = tool.Max(tmp,area)
		if height[l] <= height[r] {
			l++
		}else {
			r--
		}
	}
	return area
}


func main()  {
	nums := [] int{1,8,6,2,5,4,8,3,7}
	area := maxArea(nums)
	fmt.Println(area)
}