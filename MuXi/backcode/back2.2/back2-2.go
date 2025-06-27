package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建两个无缓冲通道用于协程同步
	letterCh := make(chan struct{})
	numCh := make(chan struct{})

	// 启动打印字母的协程
	go func() {
		for i := 0; i < 26; i++ {
			// 等待数字协程完成后接收信号
			if i > 0 {
				<-numCh
			}
			fmt.Print(string('A' + i))
			// 通知数字协程可以打印
			letterCh <- struct{}{}
		}
	}()

	// 启动打印数字的协程
	go func() {
		for i := 0; i < 26; i++ {
			// 等待字母协程完成后接收信号
			<-letterCh
			fmt.Print(i + 1)
			// 通知字母协程可以打印
			numCh <- struct{}{}
		}
	}()

	// 初始化信号，让字母协程先执行
	letterCh <- struct{}{}

	// 等待所有协程完成（这里通过通道阻塞实现）
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for range letterCh { /* 阻塞 */ }
		wg.Done()
	}()
	go func() {
		for range numCh { /* 阻塞 */ }
		wg.Done()
	}()
	wg.Wait()
}