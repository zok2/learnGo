package main

import (
	"fmt"
	"sort"
)

func payMethodScope(num int) (a, b int) {
	min1 := 10 //最小赎回
	min2 := 10 //最小保留
	Astart := min1
	Aend := num - 10
	if num-min1 <= min2 {
		Astart = num
		Aend = num
	}
	return Astart, Aend
}

func main() {
	var value int
	 var intervals[][]int
	for  i := 0; i<3; i++   {
		fmt.Scanln(&value)
		a1, b1 := payMethodScope(value)
		Scope := []int{a1,b1}
		intervals = append(intervals,Scope)
	}
	intervals = merge(intervals)
	fmt.Println(intervals)


}
func merge(intervals [][]int) [][]int {
	l := len(intervals)
	if l <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < l-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i][1] < intervals[i+1][1] {
				intervals[i][1] = intervals[i+1][1]
			}
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
			l--
		}
	}
	return intervals
}
