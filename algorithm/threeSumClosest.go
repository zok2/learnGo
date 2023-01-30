package main

import (
	"fmt"
	"math"
	"sort"
)
import "../tool"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	best := math.MaxInt32
	update := func(cur int) {
		if tool.Abs(cur-target) < tool.Abs(best-target) {
			best = cur
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				k0 := k - 1
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				j0 := j + 1
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}

		}
	}
	return best
}
func main() {
	nums := []int{-1, 2, 1, -4}
	best := threeSumClosest(nums, 1)
	fmt.Println(best)
}
