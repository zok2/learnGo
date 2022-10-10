package main

import "fmt"

func Bubble(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func Selection(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		tmp := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = tmp
	}
	return arr
}
func Insertion(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := range arr{
		preInde := i-1
		current := arr[i]
		for preInde >= 0 && arr[preInde] > current {
			arr[preInde+1] = arr[preInde]
			preInde-=1
		}
		arr[preInde+1] = current
	}
	return arr
}

func main() {
	var arr = []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, 70}
	BubbleArr := Bubble(arr)
	fmt.Println(BubbleArr)
	SelectionArr := Selection(arr)
	fmt.Println(SelectionArr)
	InsertionArr := Insertion(arr)
	fmt.Println(InsertionArr)
}
