package main

import (
	"fmt"
	"math/rand"
	"time"
)

type item struct {
	id int
	time int64
}

type result struct {
	*item
	sum int64
}

var  itemChan chan *item
var resultChan chan *result
func producer(ch  chan  *item)  {
	var id int
	for  {
		id++
		time := rand.Int63()
		tmp := &item{
			id: id,
			time : time,
		}
		ch <- tmp
	}
}
func calc(num int64) int64 {
	var sum int64
	for num>0 {
		sum = sum + num%10
		num = num/10
	}
	time.Sleep(1000*1000)
	return sum
}
func consumer(ch chan *item, resultChan chan *result)  {
	for tmp := range ch{
		sum := calc(tmp.time)
		retObj := &result{
			item: tmp,
			sum: sum,
		}
		resultChan <- retObj
		fmt.Println("来新任务了",tmp.id,time.Now().Format("2006-1-2 15:4:5"))
		time.Sleep(time.Second)
	}

}
func printResult(resultChan chan *result)  {
	for ret := range resultChan{
		fmt.Printf("ID：%v,time:%v,sum:%v,time:%v\n",ret.item.id,ret.item.time,ret.sum,time.Now().Format("2006-1-2 15:4:5"))
	}
}
func startWorker(n int ,ch chan *item,resultChan chan *result)  {
	for i := 0;i<n;i++ {
		go consumer(ch,resultChan)
	}
}
func main()  {
	itemChan = make(chan *item,100)
	resultChan = make(chan *result,100)
	go producer(itemChan)
	startWorker(20,itemChan,resultChan)
	printResult(resultChan)
}
