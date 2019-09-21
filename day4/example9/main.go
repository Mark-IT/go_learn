package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1, 8, 38, 2, 348, 484}	 // 数组是值类型，必须转成切片再sort才能完成排序
	sort.Ints(a[:])

	fmt.Println(a)
}

func testStrings() {
	var a = [...]string{"abc", "efg", "b", "A", "eeee"}		// 字符串是比较的ASCII值
	sort.Strings(a[:])

	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{2.3, 0.8, 28.2, 392342.2, 0.6}
	sort.Float64s(a[:])

	fmt.Println(a)
}

func testIntSearch() {
	var a = [...]int{1, 8, 38, 2, 348, 484}

	index := sort.SearchInts(a[:], 348)	// sort.SearchInts 查询的是排序后的值的索引
	fmt.Println(index)
	sort.Ints(a[:])
	fmt.Println(a)
}

func main() {
	//testIntSort()
	//testStrings()
	//testFloat()
	testIntSearch()
}
