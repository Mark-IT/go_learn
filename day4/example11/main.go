package main

import (
	"fmt"
	"sort"
)

func testMapSort() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
		//fmt.Println(k, v)
	}

	// map是无序的，要对map进行排序 ，即对key排序后遍历
	sort.Ints(keys)

	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func testMapKeyValExchange() { // 对map的key、val进行交换
	var a map[string]int
	var b map[int]string

	a = make(map[string]int, 5)
	b = make(map[int]string, 5)

	a["abc"] = 101
	a["efg"] = 10

	for k, v := range a {
		b[v] = k
	}

	fmt.Println(b)
}

func main() {
	testMapSort()
	testMapKeyValExchange()
}
