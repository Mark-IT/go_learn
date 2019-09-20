package main

import "fmt"

/*
defer作用：
1. 当函数返回时，执行defer语句。因此，可以用来做资源清理
2. 多个defer语句，按先进后出的方式执行
3. defer语句中的变量，在defer声明时就决定了。


用途：
1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放

 */

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

func main() {
	a()
	f()
}
