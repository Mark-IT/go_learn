package main

import (
	"errors"
	"fmt"
)

/* 内置函数
1. close：主要用来关闭channel
2. len：用来求长度，比如string、array、slice、map、channel
3. new：用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
4. make：用来分配内存，主要用来分配引用类型，比如chan、map、slice
5. append：用来追加元素到数组、slice中
6. panic和recover：用来做错误处理
 */

func testNew() {

	j := new(int)
	*j = 100        // 赋值
	fmt.Println(j)  // 地址
	fmt.Println(*j) // 值
}

func testAppend() {
	var a []int
	a = append(a, 10, 20, 30)
	fmt.Println(a)
	a = append(a, a...) // slice... 即展开slice
	fmt.Println(a)
}

func initConfig() (err error) {
	return errors.New("init config  failed")
}

func testPanicRecover() {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()

	fmt.Println("a")

	err := initConfig()
	//fmt.Println(err)
	if err != nil {
		panic(err) // 主动抛出异常
	}
	fmt.Println("b")
	fmt.Println("f")

	return
}

func main() {
	testNew()
	testAppend()
	testPanicRecover()
}
