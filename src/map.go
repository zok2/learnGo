package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var studentMap map[string]int
	println(studentMap == nil)
	studentMap = make(map[string]int, 8)
	studentMap["age"] = 18
	studentMap["results"] = 100
	for key, val := range studentMap {
		fmt.Printf(" %s => %d \n", key, val)
	}
	fmt.Printf(" len : %d ", len(studentMap))

	// slice [] = map
	studentSlice := make([]map[string]string, 8) //初始化切片
	println(studentSlice[0] == nil)
	for key, val := range studentSlice {
		if val == nil {
			val = make(map[string]string, 3)
			val["name"] = fmt.Sprint("张", key)
			val["age"] = "string(key)"
			val["results"] = fmt.Sprint(key)
			studentSlice[key] = val
			//println(val["name"],val["age"],val["results"])
		}
		if key%2 == 0 {
			delete(studentSlice[key], "age")
		}
		fmt.Println(studentSlice[key])
	}
	var resultMap = make(map[string][]int, 8)
	v, ok := resultMap["zhangsan"]
	if ok {
		fmt.Println(v)
	} else {
		resultMap["zhangsan"] = make([]int, 5)
		resultMap["zhangsan"][0] = rand.Intn(150)
		resultMap["zhangsan"][1] = rand.Intn(150)
		resultMap["zhangsan"][2] = rand.Intn(150)
		resultMap["zhangsan"][3] = rand.Intn(150)
		resultMap["zhangsan"][4] = rand.Intn(150)
		fmt.Println(resultMap)
	}


	//统计一个字符串在中单词出现的次数
	s := "how do you do"
	words := strings.Split(s," ")
	var wordsCount = make(map[string]int,4)
	for _,word :=range words{
		wordsCount[word]+=1
	}
	fmt.Println(wordsCount)
}
