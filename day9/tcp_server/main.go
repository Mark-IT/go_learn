package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server...")
	listen, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		fmt.Println("listen failed ,err:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed ,err:%v\n", err)
			continue
		}

		go process(conn)

	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			return
		}
		str := string(buf[:n])
		fmt.Printf("revc from client,data:%v\n", str)
	}
}
