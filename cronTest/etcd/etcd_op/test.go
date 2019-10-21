package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// kv.Do(op)

// kv.Put
// kv.Get
// kv.Delete

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putOp  clientv3.Op
		getOp  clientv3.Op
		opResp clientv3.OpResponse
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)

	// 创建op: operation etcd对操作的抽象
	putOp = clientv3.OpPut("/cron/jobs/job8", "op_test")

	// 执行Op
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println("error :", err)
		return
	}

	fmt.Println("写入Revision：", opResp.Put().Header.Revision)

	// 创建op
	getOp = clientv3.OpGet("/cron/jobs/job8")

	// 执行op
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println("error :", err)
		return
	}

	fmt.Println("数据Revision：", opResp.Get().Kvs[0].ModRevision)
	fmt.Println("数据value: ", string(opResp.Get().Kvs[0].Value))
}
