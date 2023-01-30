package main

import "fmt"

func longestStreak(nums []int) int {
	numMap := map[int]bool{}
	for _, num := range nums {
		numMap[num] = true
	}
	length := 0
	for num := range numMap {
		if !numMap[num-1] {
			currentNum := num
			currentStreak := 1
			for numMap[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if length < currentStreak {
				length = currentStreak
			}
		}
	}
	return length
}

func main() {
	//nums := []int{7, 54, 3, 568, 4, 6, 8, 5}
	//nums := []int{7, 54, 3, 568, 1, 6, 7567, 5}
	nums := []int{7, 54, 3, 568, 4, 6, 7567, 5}
	l := longestStreak(nums)
	fmt.Println(l)
}
