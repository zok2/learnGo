package main

import "fmt"

//1. 两数之和
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//你可以按任意顺序返回答案。

//方法一枚举
func towSumEnum(nums []int,target int) []int {
	for i,x :=range nums{
		for j := i+1;j<len(nums);j++ {
			if x+nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return nil
}
//方法二hash table
func towSumHashTable(nums []int,target int) []int {
	hashTable := map[int]int{}
	for i,x := range nums{
		if p,ok := hashTable[target-x];ok {
			return []int{p,i}
		}
		hashTable[x] = i
	}
	return nil
}
func main(){
	nums := [] int{1,3,5,7,8,0,10,-2,-1,-9}
	target := 9
	results1 := towSumEnum(nums,target)
	results2 := towSumHashTable(nums,target)
	fmt.Println(results1)
	fmt.Println(results2)
}