package main

import (
	"fmt"
)

func main() {
	stringsArrays := [][]string{
		{"apple", "appetizer", "apprehend", "app"},
		{"blueberry", "blackberry", "blossom"},
		{"cat", "caterpillar", "car"},
	}
	commonPrefixes := findCommonPrefixes(stringsArrays)
	fmt.Printf("Common Prefixes: %v\n", commonPrefixes)
}

func findCommonPrefixes(stringsArrays [][]string) []string {
	// 如果字符串数组为空，返回空切片
	if len(stringsArrays) == 0 {
		return []string{}
	}

	// 初始化公共前缀为第一个字符串数组
	commonPrefixes := stringsArrays[0]

	// 遍历其他字符串数组，找到它们的公共前缀
	for _, stringsArray := range stringsArrays[1:] {
		commonPrefixes = findCommonPrefix(commonPrefixes, stringsArray)
	}

	return commonPrefixes
}

func findCommonPrefix(strings1 []string, strings2 []string) []string {
	var commonPrefixes []string

	// 遍历两个字符串数组，找到它们的公共前缀
	for _, str1 := range strings1 {
		for _, str2 := range strings2 {
			commonPrefix := compareAndTrim(str1, str2)
			if commonPrefix != "" {
				commonPrefixes = append(commonPrefixes, commonPrefix)
			}
		}
	}

	return commonPrefixes
}

func compareAndTrim(str1, str2 string) string {
	minLen := min(len(str1), len(str2))
	commonPrefix := ""

	for i := 0; i < minLen; i++ {
		if str1[i] == str2[i] {
			commonPrefix += string(str1[i])
		} else {
			break
		}
	}

	return commonPrefix
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
