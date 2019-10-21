package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	for {
		_, err = cli.Put(context.Background(),
			"C:/Program Files/Git/logagent/con1",
			`{"one":"123", "test":1}`)
		if err != nil {
			fmt.Println("error :", err)
			return
		}
		fmt.Println("connect succ")
	}



	//for {
	//	// watch key 监听节点
	//	rch := cli.Watch(context.Background(), "C:/Program Files/Git/logagent/con1")
	//	for wresp := range rch {
	//		for _, ev := range wresp.Events {
	//			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	//			var tmp map[string]interface{}
	//			err = json.Unmarshal(ev.Kv.Value, &tmp)
	//			fmt.Println(tmp, err)
	//		}
	//	}
	//}
}
