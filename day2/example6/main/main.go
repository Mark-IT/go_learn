package main

import (
	"fmt"
	"time"
)

/*定义两个常量Man=1和Female=2，获取当前时间的秒数，如果能被Female整除，则
在终端打印female，否则打印man*/

const (
	Man    = 1
	Female = 2
)

func main() {
	for {
		second := time.Now().Unix()
		if second%Female == 0 {
			fmt.Println("Female")
		} else {
			fmt.Println("Man")
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
