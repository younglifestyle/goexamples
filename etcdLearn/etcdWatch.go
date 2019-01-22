package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"mps.longsys.com:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	cli.Put(context.Background(), "/logagent/conf/", `{"one":1, "test":1}`)
	for {
		// watch key 监听节点
		rch := cli.Watch(context.Background(), "/logagent/conf/")
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				var tmp map[string]interface{}
				err = json.Unmarshal(ev.Kv.Value, &tmp)
				fmt.Println(tmp, err)
			}
		}
	}
}
