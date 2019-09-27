package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start udp server...")

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 50002,
	})

	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}

	for {
		data := make([]byte, 1024)
		count, addr, err := listen.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("read udp failed ,err :%v\n", err)
			continue
		}
		fmt.Printf("data:%s addr:%v count:%d \n", string(data[0:count]), addr, count)

		_, err = listen.WriteToUDP([]byte("hello client"), addr)
		if err != nil {
			fmt.Printf("write udp failed ,err:%v\n", err)

			continue
		}

	}

}
