package main

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for  low<=high{
		mid := (low+high)/2+low
		num := nums[mid]
		if num == target {
			return mid
		}else if num > target {
			high = mid - 1
		}else {
			low = mid+1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	binarySearch(nums, target)
}

