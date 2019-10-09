package main

import (
	"time"
)

func main() {
	initRedis("localhost:6379", "pgc123456",16, 1024, time.Second*300)
	initUserMgr()
	runServer("0.0.0.0:10000")
}
