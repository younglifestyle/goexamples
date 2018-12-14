package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fagongzi/gateway/pkg/pb/metapb"

	"github.com/fagongzi/gateway/pkg/client"
	"github.com/gin-gonic/gin"
)

var clusterA uint64

func main() {
	router := gin.Default()

	router.GET("/hello/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, "yes, this is hello")
	})
	router.POST("/hello/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, "yes, this is hello")
	})

	// 第一步： 创建Cluster，类似于服务分类
	if err := createCluster1(); err != nil {
		log.Println("1 error,", err)
		return
	}

	// 第二步： 对应真正的业务服务器
	err := createServer1()
	if err != nil {
		log.Println("2 error,", err)
		return
	}

	// 第三步： 创建API，该API会被转发到ClusterA
	if err := createAPI1(); err != nil {
		log.Println("3 error,", err)
		return
	}

	router.Run(":8068")
}

func createCluster1() error {
	c, err := getClient()
	if err != nil {
		return err
	}

	clusterA, err = c.NewClusterBuilder().Name("cluster-A").
		Loadbalance(metapb.RoundRobin).Commit()
	if err != nil {
		return err
	}

	return nil
}

func createServer1() error {
	c, err := getClient()
	if err != nil {
		return err
	}

	sb := c.NewServerBuilder()
	// 必选项
	sb.Addr("127.0.0.1:8068").HTTPBackend().MaxQPS(100)

	id, err := sb.Commit()
	if err != nil {
		return err
	}

	// 把这个server加入到cluster A
	c.AddBind(clusterA, id)
	return nil
}

func createAPI1() error {
	c, err := getClient()
	if err != nil {
		return err
	}

	sb := c.NewAPIBuilder()
	// 必选项
	sb.Name("用户API")
	// 设置URL规则，匹配所有开头为/api/user的请求
	sb.MatchURLPattern("/hello/(.+)")
	// 匹配GET请求
	sb.MatchMethod("GET")
	// 添加tag
	sb.AddTag("webhook", "true")
	// 匹配所有请求
	//sb.MatchMethod("*")
	// 不启动
	//sb.Down()
	// 启用
	sb.UP()
	// 分发到Cluster A
	sb.AddDispatchNode(clusterA)

	id, err := sb.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("api id is: %d", id)
	return nil
}

// 如果你的api server使用了"--discovery"参数启动
func getClientWithDiscovery() (client.Client, error) {
	return client.NewClientWithEtcdDiscovery("/services",
		time.Second*10,
		"127.0.0.1:2379")
}

// 如果你的api server没有使用"--discovery"参数启动
func getClient() (client.Client, error) {
	return client.NewClient(time.Second*10,
		"127.0.0.1:9091")
}
