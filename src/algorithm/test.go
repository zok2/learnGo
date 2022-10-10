package main
package main

import "math"

func minDifference(nums []int, queries [][]int) []int {
	var min func(a, b int) int
	min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	n := len(nums)
	prefix := make([][]int, n+1)
	prefix[0] = make([]int, 101)
	for i := 0; i < n; i++ {
		prefix[i+1] = make([]int, 101)
		copy(prefix[i+1], prefix[i])
		prefix[i+1][nums[i]]++
	}
	q := len(queries)
	res := []int{}
	for i := 0; i < q; i++ {
		left, right := queries[i][0], queries[i][1]
		last, best := 0, math.MaxInt32
		for j := 1; j <= 100; j++ {
			if prefix[left][j] != prefix[right+1][j] {
				if last != 0 {
					best = min(best, j-last)
				}
				last = j
			}
		}
		if best == math.MaxInt32 {
			best = -1
		}
		res = append(res, best)
	}
	return res
}
