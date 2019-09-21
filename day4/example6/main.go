package main

import "fmt"

func fab(n int) {
	var a []uint64
	a = make([]uint64, n)
	a[0] = 1
	a[1] = 1

	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	for _, v := range a {
		fmt.Println(v)
	}
}

func testArray() { //数组初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{1, 2, 3, 4, 5, 6, 7, 8} // 根据初始化的值，自动定义长度
	var a3 = [...]int{1: 300, 3: 200}         // 指定位置赋值初始化,长度由最后一个索引位置决定
	var a4 = [...] string{1: "hello", 3: "world"}

	fmt.Println("a", a)
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a3", a3)
	fmt.Println("a4", a4)
}

func testArray2() {
	var a [2][5]int = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	for row, v := range a {
		for col, v1 := range v {
			fmt.Printf("(%d,%d)=%d", row, col, v1)
		}
		fmt.Println()
	}
}

func main() {
	fab(10)
	testArray()
	testArray2()
}
