// 题目：要求开启50个协程去并发地将一个变量从0增加到5000,也就是说每个协程需要加给这个数加100,但是 每个协程每次只能给这个变量加1
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	mu    sync.Mutex
	wg    sync.WaitGroup
)

func add() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

func main() {
	// 初始化WaitGroup，等待50个协程
	wg.Add(50)
	
	// 启动50个协程并发加数
	for i := 0; i < 50; i++ {
		go add()
	}
	
	// 等待所有协程完成
	wg.Wait()
	fmt.Println("最终计数:", count)
	time.Sleep(time.Second)
}