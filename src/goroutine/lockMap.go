package main

import (
"fmt"
"strconv"
"sync"
)

var d2 = make(map[string]int)
var w2 sync.WaitGroup

var lo sync.RWMutex

func getValue(key string) int {
	lo.RLock()
	defer lo.RUnlock()
	return d2[key]
}


func setValue(key string,value int) {
	lo.Lock()
	defer lo.Unlock()
	d2[key]=value
}

type slice struct {
	
}

func main(){
	for i :=0;i<3000;i++{
		w2.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			setValue(key,n)
			fmt.Printf("k=:%v,v:=%v\n",key,getValue(key))
			w2.Done()
		}(i)
	}
	w2.Wait()
}