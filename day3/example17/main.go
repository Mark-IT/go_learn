package main

func main() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE // 此处goto label 相当于实现了一个for循环

}
