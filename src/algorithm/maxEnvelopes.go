package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})
	dp := make([]int, len(envelopes))
	for i, _ := range dp {
		dp[i] = 1
	}
	MM := 0
	for i := 0; i < len(dp); i++ {

		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)

			}
		}
		MM = max(MM, dp[i])
	}
	return MM
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	boxes := [][]int{{56, 47}, {63, 42}, {66, 71}, {25, 39}}
	t := maxEnvelopes(boxes)
	fmt.Println(t)
}
