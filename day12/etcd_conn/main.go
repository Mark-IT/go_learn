package main

import (
	"fmt"
	etcdClient "go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints:   []string{"localhost:2379", "localhost:12379", "localhost:22379"}, //  即使有服务挂了，也没关系，自动做健康检查，容错处理，这里也可以填域名
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect to etcd failed, err:", err)
		return
	}

	fmt.Println("connect to etcd success")
	defer cli.Close()
}
