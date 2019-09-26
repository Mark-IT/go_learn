package main

import (
	"fmt"
	"time"
)

/*
 select 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。
在执行select语句的时候，运行时系统会自上而下地判断每个case中的发送或接收操作是否可以被立即执行
立即执行：意思是当前Goroutine不会因此操作而被阻塞，还需要依据通道的具体特性(缓存或非缓存


每个case语句里必须是一个IO操作
所有channel表达式都会被求值、所有被发送的表达式都会被求值
如果任意某个case可以进行，它就执行(其他被忽略)。
如果有多个case都可以运行，Select会随机公平地选出一个执行(其他不会执行)。
如果有default子句，case不满足条件时执行该语句。
如果没有default字句，select将阻塞，直到某个case可以运行；Go不会重新对channel或值进行求值。

select分支选择规则
所有跟在case关键字右边的发送语句或接收语句中的通道表达式和元素表达式都会先被求值。无论它们所在的case是否有可能被选择都会这样。
求值顺序：自上而下、从左到右

 */
func main() {
	var ch chan int
	ch = make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		var i int
		for {
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i * i
			time.Sleep(time.Second)
			i++
		}
	}()
	for {
		select { // select + time.After 实现超时控制
		case v := <-ch:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <-time.After(time.Second): // 超过1秒没获取到值，则执行
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}
