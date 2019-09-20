package main

import (
	"fmt"
	"time"
)

// 写一个程序，统计一段代码的执行耗时，单位精确到微秒

func test() {
	// 等待 10毫秒 => 即 10000微秒
	time.Sleep(time.Millisecond * 10)
}

func main() {
	startTime := time.Now().UnixNano() //	返回时间戳，包含纳秒
	test()
	endTime := time.Now().UnixNano()
	fmt.Println("cost  us:", (endTime-startTime)/1000)
}
