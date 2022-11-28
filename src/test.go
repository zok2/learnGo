package main

import "fmt"

//func isUnique(str string) bool {
//	md := make(map[rune]int)
//	for _, v := range str {
//		md[v]++
//		if md[v] > 1 {
//			return false
//		}
//	}
//	return true
//}
//func main()  {
//	f := isUnique("abc")
//	fmt.Println(f)
//}
//
//func bubbleSort(arr []int) []int {
//	length := len(arr)
//	for i := 0; i < length; i++ {
//		for j := 0; j < length-1-i; j++ {
//			if arr[j] > arr[j+1] {
//				arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//	return arr
//}
////插入
//func insertionSort(arr []int) []int {
//	for i := range arr {
//		preIndex := i - 1
//		current := arr[i]
//		for preIndex >= 0 && arr[preIndex] > current {
//			arr[preIndex+1] = arr[preIndex]
//			preIndex -= 1
//		}
//		arr[preIndex+1] = current
//	}
//	return arr
//}
func hello(i int) {
	fmt.Println(i)
}
func main() {
	primeNums := [6]int{2, 3, 5, 7, 11, 13}
	slice1 := primeNums[1:4]
	fmt.Println(slice1)
}