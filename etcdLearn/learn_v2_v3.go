package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"

	"log"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func getIP() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func ScanEtcd() (string, error) {
	// return "http://mps.longsys.com:2379", nil
	server := os.Getenv("ETCD_SERVER_URL")
	if server == "" {
		server = "http://mps.longsys.com:2379"
	}
	return server, nil
}

/* return:
common config viper handle
private config viper handle
error
*/
func InitEtcd() (*viper.Viper, error) {
	var common_viper = viper.New()

	url, err := ScanEtcd()
	if err != nil {
		fmt.Println("Etcd Service not found")
		return nil, err
	}

	//ip := getIP()

	common_path := fmt.Sprintf("/one")

	fmt.Printf("ETCD Server Url: %s\n", url)
	fmt.Printf("Common Config Path: %s\n", common_path)

	err = common_viper.AddRemoteProvider("etcd", url, common_path)
	if err != nil {
		fmt.Printf("Read common AddRemoteProvider error: %s\n", err.Error())
		return nil, err
	}

	common_viper.SetConfigType("json")

	err = common_viper.ReadRemoteConfig()
	if err != nil {
		fmt.Printf("Read common RemoteConfig error: %s\n", err.Error())
		return nil, err
	}

	return common_viper, nil
}

func WatchEtcd(runtime_viper *viper.Viper, interval time.Duration) {
	go func() {
		for {
			time.Sleep(time.Second * interval) // delay after each request

			// currently, only tested with etcd support
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				fmt.Printf("unable to read remote config: %v\n", err)
				continue
			}
		}
	}()
}

func connectEtcd() {
	/*
		DialTimeout time.Duration `json:"dial-timeout"`
		Endpoints []string `json:"endpoints"`
	*/
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

	//设置1秒超时，访问etcd有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//操作etcd
	_, err = cli.Put(ctx, "/logagent/conf/", "sample_value")
	//操作完毕，取消etcd
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
}

func main() {

	connectEtcd()

	viperCommon, err := InitEtcd()
	if err != nil {
		log.Fatal("error init etcd")
	}

	WatchEtcd(viperCommon, 5)

	time.Sleep(6)

	log.Println(viperCommon.Get("toolmanager"))
}
