package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("read data")
}

func (f *File) Write() {
	fmt.Println("write data")
}

func Test(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	var f *File
	var b interface{}
	b = f
	Test(f)	// File实现了 Read 和 Write方法，所以可以直接使用Test
	v, ok := b.(ReadWriter)	// 判断b是否是ReadWriter
	fmt.Println(v, ok)

}
