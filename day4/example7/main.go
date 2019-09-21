package main

import "fmt"

/*
1. 切片：切片是数组的一个引用，因此切片是引用类型
2. 切片的长度可以改变，因此，切片是一个可变的数组
3. 切片遍历方式和数组一样，可以用len()求长度
4. cap可以求出slice最大的容量，0 <= len(slice) <= (array)，其中array是slice引用的数组
5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int



 */

type slice struct {
	// 自定义模仿一个切片类型
	ptr *[100]int
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	s.ptr[1] = 1000
}

func testSlice2() {
	var s1 slice
	s1 = make1(s1, 10)

	s1.ptr[0] = 100
	modify(s1)

	fmt.Println(s1.ptr)
}

func testSlice() {
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	slice = arr[:]
	slice = slice[1:]
	slice = slice[:len(slice)-1] // 去掉切片最后一个元素
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}

func modify1(a []int) { // 切片是引用类型，可以直接修改值
	a[1] = 1000
}

func testSlice3() {
	var b []int = []int{1, 2, 3, 4}
	modify1(b)
	fmt.Println(b)
}

func testSlice4() {
	var a = [10]int{1, 2, 3, 4}

	//b := a[1:5]
	b := a[:]
	fmt.Printf("%p\n", b) // b是一个切片，其内存地址指向数组切的起始位置
	fmt.Printf("%p\n", &a[1])
}

func testSlice5() { // 通过 make创建切片,第一个参数是切片类型，第二个是len,第三个是cap
	var slice []int = make([]int, 5)
	slice1 := make([]int, 5)
	slice2 := make([]int, 5, 10)
	fmt.Printf("slice:%v,len:%d,cap:%d\n", slice, len(slice), cap(slice))
	fmt.Printf("slice1:%v,len:%d,cap:%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2:%v,len:%d,cap:%d\n", slice2, len(slice2), cap(slice2))
}
func main() {
	//testSlice()
	//testSlice2()
	//testSlice3()
	//testSlice4()
	testSlice5()
}
