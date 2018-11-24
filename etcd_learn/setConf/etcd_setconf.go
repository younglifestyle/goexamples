package main

import (
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"

	"github.com/gin-gonic/gin/json"
)

//context是一个在go中时常用到的程序包，google官方开发。
//特别常见的一个应用场景是由一个请求衍生出的各个goroutine之间需要满足一定的约束关系，
//以实现一些诸如有效期，中止routine树，传递请求全局变量之类的功能。
//使用context实现上下文功能约定需要在
//你的方法的传入参数的第一个传入一个context.Context类型的变量。

const (
	EtcdKey = "/oldboy/backend/logagent/config/192.168.14.4"
)

type LogConf struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func setLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"36.111.184.221:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect success...")
	defer cli.Close()

	logConfArr := []LogConf{
		LogConf{
			Path:  "D:/project/nginx/logs/access.log",
			Topic: "nginx_log",
		},
		LogConf{
			Path:  "D:/project/nginx/logs/error.log",
			Topic: "nginx_log_err",
		},
	}

	marshalData, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed :", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(marshalData))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Println("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	setLogConfToEtcd()
}
