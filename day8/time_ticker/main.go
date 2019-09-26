package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	t := time.NewTicker(time.Second)	//初始化断续器,间隔2s
	go func() {
		for {

			select {
			case <-t.C:
				fmt.Println("timeout")
			}

		}
	}()

	time.Sleep(time.Second * 5)	//阻塞，则执行次数为sleep的休眠时间/ticker的时间
	t.Stop()	// 用来停止定时器
}
