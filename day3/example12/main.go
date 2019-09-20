package main

import "fmt"

func main() {

	//a:= 0
	//switch  {
	switch a := 1; { // 合并写时，这里必须加分号
	case a == 0:
		fmt.Println("0")
	case a >= 1 && a <= 3: // a 为1或2时，都执行下面的代码
		fmt.Println("1<=a<=3")
	default:
		fmt.Println(" a>3")

	}
}
