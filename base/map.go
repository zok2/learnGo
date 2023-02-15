package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//声明变量，map 默认是 nil向一个nil值的map存入元素将导致一个panic异常：panic: assignment to entry in nil map
	var studentMap map[string]int
	println("studentMap :", studentMap == nil)
	studentMap = make(map[string]int)
	studentMap["age"] = 18
	studentMap["results"] = 100
	fmt.Println("studentMap :", studentMap, "len :", len(studentMap))

	// 声明变量时，进行初始化
	var cityMap = map[string]string{"name": "北京", "code": "001"}
	fmt.Println("cityMap :", cityMap, "len :", len(cityMap))

	//删除元素
	delete(studentMap, "age")
	//map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作：
	//_ = &ages["bob"] // compile error: cannot take address of map element
	//禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，
	//从而可能导致之前的地址无效。

	//遍历map元素
	for k, v := range studentMap {
		fmt.Println("key ：", k, " => ", "val : ", v)
	}

	// slice [] = map
	studentSlice := make([]map[string]string, 8) //初始化切片
	for key, val := range studentSlice {
		if val == nil {
			val = make(map[string]string, 3)
			val["name"] = fmt.Sprint("张", key)
			val["age"] = strconv.Itoa(key)
			val["results"] = fmt.Sprint(key)
			studentSlice[key] = val
			//println(val["name"],val["age"],val["results"])
		}
		if key%2 == 0 {
			delete(studentSlice[key], "age")
		}
		fmt.Println(studentSlice[key])
	}

	//统计一个字符串在中单词出现的次数
	s := "how do you do"
	words := strings.Split(s, " ")
	var wordsCount = make(map[string]int, 4)
	for _, word := range words {
		wordsCount[word] += 1
	}
	fmt.Println("统计字符 wordsCount 出现的次数", wordsCount)
}
