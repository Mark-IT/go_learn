package main

import (
	"fmt"
	"strconv"
)

/*
打印出100-999中所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字
立方和等于该数本身。例如：153 是一个“水仙花数”，因为 153=1 的三次
方＋5 的三次方＋3 的三次方
 */

func isNarcissus(n int) bool {
	var i, j, k int
	i = n % 10         // 个位
	j = (n / 10) % 10  // 十位
	k = (n / 100) % 10 // 百位
	sum := i*i*i + j*j*j + k*k*k
	return sum == n
}

func isNarcissus1(str string) bool {

	result := 0
	for i := 0; i < len(str); i++ {
		num := int(str[i] - '0') // 通过ASCII码值计算
		result += num * num * num
	}

	number, err := strconv.Atoi(str)

	if err != nil {
		fmt.Printf("can not convert %s to int", str)
		return false
	}
	if result == number {
		return true
	}

	return false
}

func main() {
	var n int // 区间起始值
	var m int // 区间终点值
	fmt.Scanf("%d%d", &n, &m)
	for i := n; i <= m; i++ {
		if isNarcissus(i) == true {
		//if isNarcissus1(strconv.Itoa(i)) == true {
			fmt.Printf("[%d] is narcissus number \n", i)
			continue
		}
	}
}
