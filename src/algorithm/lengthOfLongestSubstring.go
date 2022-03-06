package main

import (
	"fmt"
)
//滑动窗口  通过位移右指针验证，将字符value ++   不等0 跳出循环 记录最大值，位移左指针，删除刚才重复出现的值。
func lengthOfLongestSubstring(s string) ( ans int) {
	m := map[byte]int{}
	n :=len(s)
	rk,ans :=-1,0
	for i := 0; i < n ;i++ {
		if i != 0  {
			delete(m,s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0  {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans,rk - i +1)
	}
	return
}
func max(int1 int,int2 int) int {
	if int1 > int2 {
		return int1
	}
	return int2
}
func main()  {
	s :="abcabcbb"
	leng := lengthOfLongestSubstring(s)
	fmt.Println(leng)
}
