package main

// map  线程安全
import (
	"fmt"
	"strconv"
	"sync"
)

var d = make(map[string]int)
var w sync.WaitGroup

func get(key string) int {
	return d[key]
}


func set(key string,value int) {
	d[key]=value
}



func main(){
	for i :=0;i<300;i++{
		w.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key,n)
			fmt.Printf("k=:%v,v:=%v\n",key,get(key))
			w.Done()
		}(i)
	}
	w.Wait()
}