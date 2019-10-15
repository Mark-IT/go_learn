package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{} //  一般用于主线程为了等待goroutine都运行完毕

	for i := 0; i < 10; i++ {
		wg.Add(1) // Add(delta)是将计数器加减delta，Done()就是调用Add(-1)。
		go calc(&wg, i)
	}

	wg.Wait() // wait() 会阻塞代码的运行，直到计数器地值减为0。
	fmt.Println("all goroutine finish")
}

// 一定要通过指针传值，不然进程会进入死锁状态
func calc(w *sync.WaitGroup, i int) {

	fmt.Println("calc:", i)
	time.Sleep(time.Second)
	w.Done() // Done() 每次把计数器-1
}
