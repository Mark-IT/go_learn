package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:50001")

	if err != nil {
		fmt.Printf("dial failed,err :%v\n", err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		data, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from consolo failed , err:%v\n", err)
			break
		}
		data = strings.TrimSpace(data) // 去掉换行符

		if data == "Q" {
			break
		}
		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Println("write failed ,err :%v\n ", err)
			return
		}
	}
}
