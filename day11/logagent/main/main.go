package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go_learn/day11/logagent/kafka"
	"go_learn/day11/logagent/tailf"
	//"time"
)

func main() {
	filename := "./logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}

	logs.Debug("load conf success, config:%v", appConfig)

	collectConf, err := initEtcd(appConfig.etcdAddr, appConfig.etcdKey) // 从etcd中读取配置
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}
	logs.Debug("initialize etcd success")

	err = tailf.InitTail(collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	logs.Debug("initialize tailf success")
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	logs.Debug("initialize all succ")
	/*
		go func() {
			var count int
			for {
				count++
				logs.Debug("test for logger %d", count)
				time.Sleep(time.Millisecond * 1000)
			}
		}()*/
	err = serverRun()
	if err != nil {
		logs.Error("serverRUn failed, err:%v", err)
		return
	}

	logs.Info("program exited")
}
