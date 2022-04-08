package main
import (
	"../tool"
	"fmt"
)
//滑动窗口  通过位移右指针验证，将字符value ++   不等0 跳出循环 记录最大值，位移左指针，删除刚才重复出现的值。
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}

		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = tool.Max(ans, rk-i+1)
	}
	return ans
}


func main() {
	str := "pwwkew"
	count := lengthOfLongestSubstring(str)
	fmt.Println([]byte(str))
	fmt.Println(count)
}
