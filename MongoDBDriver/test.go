package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// mongodb日志管理
type LogMgr struct {
	client        *mongo.Client
	logCollection *mongo.Collection
}

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

// 一条日志
type LogRecords struct {
	JobName string `bson:"jobName"` // 任务名
	Command string `bson:"command"` // shell命令
	Err     string `bson:"err"`     // 脚本错误
	Content string `bson:"content"` // 脚本输出
	TimePoint
}

var (
	G_logMgr *LogMgr
)

func InitLogMgr() {

	// 建立MongoDB连接
	opts := options.Client()
	opts.ApplyURI("mongodb://172.16.9.229:27017")
	opts.SetConnectTimeout(time.Duration(
		2) * time.Millisecond)
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

func main() {
	InitLogMgr()

	// 4, 插入记录(bson)
	record := &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	oneResult, err := G_logMgr.logCollection.InsertOne(context.TODO(), record)
	fmt.Println(oneResult, err)

	records := &LogRecords{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}
	oneResult, err = G_logMgr.logCollection.InsertOne(context.TODO(), records)
	fmt.Println(oneResult, err)

	objectIDS, _ := primitive.ObjectIDFromHex("5caeaa0a170214f7a58319fe")

	var recordTmp LogRecord
	filter := bson.M{"_id": objectIDS}
	singleResult := G_logMgr.logCollection.FindOne(context.TODO(), filter)
	err = singleResult.Decode(&recordTmp)
	fmt.Println(err, recordTmp)

	// 单次查询没有找到，将会报错
	objectIDS, _ = primitive.ObjectIDFromHex("5caeaa0a170214f7a58319ff")
	filter = bson.M{"_id": objectIDS}
	singleResult = G_logMgr.logCollection.FindOne(context.TODO(), filter)
	err = singleResult.Decode(&recordTmp)

	fmt.Println()
	fmt.Println(err)
	//fmt.Println(mongo.ErrNoDocuments)

	objectIDS, _ = primitive.ObjectIDFromHex("5caeaa0a170214f7a58319ff")
	filter = bson.M{"_id": objectIDS}
	_, err = G_logMgr.logCollection.Find(context.TODO(), filter)

	fmt.Println()
	fmt.Println(err)
}
