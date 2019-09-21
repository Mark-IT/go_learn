package main

import "fmt"

// 实现一个插入排序

func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	val := a[left]	// 选取最左边的数，保存其值
	k := left // 保存其下标

	for i := left + 1; i <= right; i++ {
		if a[i] < val {		// 如果右边的值比当前选取的值小，则将其移动到被选取值的左边
			a[k] = a[i]		// 将当前比较的值移动到原选择的位置，这样就空出了其位置
			a[i] = a[k+1]
			k++	// 移动下标

		}

	}
	a[k] = val	// 将选取的值放到最终位置，此时val所在位置的左边都比其小，右边都比其大
	quickSort(a, left, k-1)
	quickSort(a, k+1, right)
}

func main() {
	b := [...]int{9, 1, 3, 6, 5, 10, 8}
	fmt.Println(b)
	quickSort(b[:], 0, len(b)-1)
	fmt.Println(b)

}
