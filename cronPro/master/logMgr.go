package master

import (
	"context"
	"goexamples/cronPro/common"
	"log"
	"time"

	"github.com/mongo-go-driver/mongo/findopt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// mongodb日志管理
type LogMgr struct {
	client        *mongo.Client
	logCollection *mongo.Collection
}

var (
	G_logMgr *LogMgr
)

func InitLogMgr() {

	// 建立MongoDB连接
	opts := options.Client()
	opts.ApplyURI(G_config.MongodbUri)
	opts.SetConnectTimeout(time.Duration(
		G_config.MongodbConnectTimeout) * time.Millisecond)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("connect mongodb is error")
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("mongodb server error")
		return
	}

	G_logMgr = &LogMgr{
		client: client,
		logCollection: client.Database("cron").
			Collection("log"),
	}

	return
}

// 查看任务日志
func (logMgr *LogMgr) ListLog(name string,
	skip, limit int) (logArr []*common.JobLog, err error) {

	// len(logArr)
	logArr = make([]*common.JobLog, 0)

	// 过滤条件
	filter := bson.M{"name": "pi"}

	sorting := map[string]int{}
	// 按照任务开始时间倒排
	sorting["startTime"] = -1

	cursor, err := logMgr.logCollection.Find(context.TODO(),
		filter, findopt.Sort(sorting), findopt.Skip(int64(skip)),
		findopt.Limit(int64(limit)))
	if err != nil {
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		jobLog := &common.JobLog{}

		if err = cursor.Decode(jobLog); err != nil {
			continue
		}

		logArr = append(logArr, jobLog)
	}
	return
}
