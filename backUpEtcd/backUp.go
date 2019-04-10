package main

import (
	"context"
	"fmt"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type EtcdStore struct {
	sync.RWMutex
	prefix    string
	rawClient *clientv3.Client
	kv        clientv3.KV
}

func main() {
	var etcdStore EtcdStore

	err := etcdStore.ConnectEtcd("www.wukoon-app.com:2379")
	if err != nil {
		fmt.Println(err)
	}

	etcdStore.SetPrefixPath("/backup/test")

	err = etcdStore.BackupTo("127.0.0.1:2379")
	if err != nil {
		fmt.Println(err)
	}

}

func (e *EtcdStore) ConnectEtcd(addr string) (err error) {
	// 初始化配置
	config := clientv3.Config{
		Endpoints:   []string{addr},  // 集群地址
		DialTimeout: time.Second * 2, // 连接超时
	}
	// 建立连接
	e.rawClient, err = clientv3.New(config)
	if err != nil {
		return err
	}

	e.kv = clientv3.NewKV(e.rawClient)

	return nil
}

func (e *EtcdStore) SetPrefixPath(path string) {
	e.prefix = path
}

func test(to string) (err error) {
	// 建立连接
	targetEtcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{to},    // 集群地址
		DialTimeout: time.Second * 2, // 连接超时
	})
	if err != nil {
		return err
	}

	defer targetEtcdClient.Close()
	kv := clientv3.NewKV(targetEtcdClient)

	// 执行OP
	opResp, err := kv.Do(context.TODO(),
		clientv3.OpGet("/backup/test/1"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(opResp.Get().Kvs[0].Value))

	return nil
}

func (e *EtcdStore) BackupTo(to string) (err error) {
	e.Lock()
	defer e.Unlock()

	// 建立连接
	targetEtcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{to},    // 集群地址
		DialTimeout: time.Second * 2, // 连接超时
	})
	if err != nil {
		return err
	}

	defer targetEtcdClient.Close()
	kv := clientv3.NewKV(targetEtcdClient)

	// get一下`prefix`目录下的所有信息
	getResp, err := e.kv.Get(context.TODO(),
		e.prefix, clientv3.WithPrefix())
	if err != nil {
		return
	}

	for _, kvpair := range getResp.Kvs {
		fmt.Println(string(kvpair.Key), string(kvpair.Value))
		// 创建Op: operation  对操作的抽象
		putOp := clientv3.OpPut(string(kvpair.Key), string(kvpair.Value))
		// 执行OP
		opResp, err := kv.Do(context.TODO(), putOp)
		if err != nil {
			fmt.Println(err, opResp.Put().Header.Revision)
			return err
		}
	}

	return nil
}

// 提取worker的IP
func ExtractPrefix(regKey, prefix string) string {
	return strings.TrimPrefix(regKey, prefix)
}

func (e *EtcdStore) BackupToSpecifyDir(toAddr, newPrefix string) (err error) {
	e.Lock()
	defer e.Unlock()

	// 建立连接
	targetEtcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{toAddr}, // 集群地址
		DialTimeout: time.Second * 2,  // 连接超时
	})
	if err != nil {
		return err
	}

	defer targetEtcdClient.Close()
	kv := clientv3.NewKV(targetEtcdClient)

	// get一下`prefix`目录下的所有信息
	getResp, err := e.kv.Get(context.TODO(),
		e.prefix, clientv3.WithPrefix())
	if err != nil {
		return
	}

	for _, kvpair := range getResp.Kvs {
		fmt.Println(string(kvpair.Key), string(kvpair.Value))

		// 重组路径
		newEtcdPath := path.Join(newPrefix, ExtractPrefix(string(kvpair.Key), e.prefix))

		// 创建Op: operation  对操作的抽象
		putOp := clientv3.OpPut(newEtcdPath, string(kvpair.Value))
		// 执行OP
		opResp, err := kv.Do(context.TODO(), putOp)
		if err != nil {
			fmt.Println(err, opResp.Put().Header.Revision)
			return err
		}
	}

	return nil
}
