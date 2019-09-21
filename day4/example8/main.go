package main

import "fmt"

func testSlice() {
	a := [...]int{1, 2, 3, 4, 5}
	s := a[1:] // 当前len和cap都是4
	fmt.Println("a:", a)
	fmt.Println("s:", s)
	fmt.Printf("before s len[%d] cap[%d]\n", len(s), cap(s))
	s[1] = 100 // 扩容前的，由于内存地址相同，所以对s的修改会提现在a上
	fmt.Printf("before s=%p a[1]=%p\n", s, &a[1])
	fmt.Println("before a:", a)
	fmt.Println("before s:", s)

	s = append(s, 6) // 添加第5个元素，超出cap,自动扩容，cap=cap*2
	fmt.Printf("after s len[%d] cap[%d]\n", len(s), cap(s))
	s = append(s, 7)
	s = append(s, 8)
	s = append(s, 9)
	s = append(s, 10)

	s[1] = 1000 // 扩容后,内存地址已发生变化，此时再对s操作，不会修改a
	fmt.Printf("after s=%p a[1]=%p\n", s, &a[1])
	fmt.Println("after a:", a)
	fmt.Println("after s:", s)

}

func testCopy() {

	var a []int = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 1, 3)
	fmt.Printf("before b:%v len:%d cap:%d \n", b, len(b), cap(b))
	copy(b, a) // 切片拷贝，将a的值，拷贝到b中，拷贝多少取决于b的len

	fmt.Printf("after b:%v len:%d cap:%d \n", b, len(b), cap(b))
	var c []int
	c = append(c, b...)
	c = append(c, 9, 9, 9)
	fmt.Println(c)
}

func testString() {
	s := "hello world" // string底层就是一个byte的数组，因此，也可以进行切片操作
	s1 := s[0:5]
	s2 := s[6:]

	fmt.Println(s1)
	fmt.Println(s2)

}

func testModifyString() {
	s := "我hello world" // 	string本身是不可变的，因此要改变string中字符，需要将其转成切片
	s1 := []rune(s)

	s1[0] = 200
	s1[1] = 128
	s1[2] = 64
	str := string(s1)
	fmt.Println(str)
}

func main() {
	//testSlice()
	testCopy()
	//testString()
	//testModifyString()
}
