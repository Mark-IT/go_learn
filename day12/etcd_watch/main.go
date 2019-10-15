package main

import (
	"context"
	"fmt"
	etcdClient "go.etcd.io/etcd/clientv3"
	"time"
)

// watch用来获取未来更改的通知。 测试时，在命令行里需要先将API version 设置为3，才能正常监控，默认为2
func main() {

	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints:   []string{"localhost:2379", "localhost:12379", "localhost:22379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	rch := cli.Watch(context.Background(), "/logagent/conf/")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
