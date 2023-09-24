package main

import (
	"fmt"
	"math"
)

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

func test(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]

		}
	}
	return dp[amount]
}
func main() {
	coins := []int{1, 2, 5}
	amount := 12
	c := test(coins, amount)
	fmt.Printf("amount:%d , count : %d", amount, c)
}
