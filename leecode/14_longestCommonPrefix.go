package main

import (
	"fmt"
)

//14. 最长公共前缀
/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：
提示：
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
*/

// 横向
func longestCommonPrefix1(strs []string) string {

	count := len(strs)
	prefix := strs[0]
	for i := 1; i < count; i++ {
		prefix = diff(prefix, strs[i])
	}
	return prefix
}

func diff(string1 string, string2 string) string {
	index := 0
	lenght := min(len(string1), len(string2))
	for index < lenght && string1[index] == string2[index] {
		index++
	}
	return string1[:index]
}

// 纵向
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {

	var strs = make([]string, 3, 3)
	strs = []string{"flower", "flow", "flight"}
	r1 := longestCommonPrefix1(strs)
	r2 := longestCommonPrefix2(strs)
	fmt.Println(r1)
	fmt.Println(r2)

}
