package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
go build 时 可以指定参数 -race ，运行编译后的程序时，可以看出是否存在竞争

go get 命令用来安装第三方包，安装路径在gopath里，跟命令执行位置无关
 */

var lock sync.Mutex     // 互斥锁
var rwLock sync.RWMutex //读写锁

func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}

	lock.Lock()
	fmt.Println(a) // 读也要加锁
	lock.Unlock()

	time.Sleep(time.Second)
}

func testRWLock() { // 读多写少使用读写锁，效率比互斥锁高的多
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			//lock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			//lock.Unlock()
			rwLock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				//lock.Lock()
				rwLock.RLock()
				time.Sleep(time.Millisecond)
				fmt.Println(a)
				rwLock.RUnlock()
				//lock.Unlock()
				atomic.AddInt32(&count, 1) // 原子操作加
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count)) // 原子操作读
}

func main() {
	//testMap()
	testRWLock()
}
