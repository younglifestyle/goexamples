package main

import (
	"fmt"
	"goexamples/kafka/logagent/kafka"
	"goexamples/kafka/logagent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
	//"fmt"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	fmt.Printf("read msg:%s, topic:%s\n", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
