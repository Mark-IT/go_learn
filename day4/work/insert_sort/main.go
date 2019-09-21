package main

import "fmt"

// 实现一个插入排序

func insertSort(a []int) {
	for i := 1; i < len(a); i++ { // 默认第一个元素自己已经有序，从第二个元素开始，跟第一个元素比较后插入
		for j := i; j > 0; j-- { // 将无序区与有序区每个元素比较
			if a[j] > a[j-1] { // 已经有序，直接跳出，判断下一个无序区元素
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
		}

	}
}

func main() {
	b := [...]int{9, 1, 3, 6, 5, 10, 8}
	fmt.Println(b)
	insertSort(b[:])
	fmt.Println(b)

}
