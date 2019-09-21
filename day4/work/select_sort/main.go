package main

import "fmt"

// 实现一个选择排序

func selectSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i	// 选择最小的数的下标
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j

			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

func main() {
	b := [...]int{9, 1, 3, 6, 5, 10, 8}
	fmt.Println(b)
	selectSort(b[:])
	fmt.Println(b)

}
