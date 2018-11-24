package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	//filename := "E:\\project\\kafka_2.12-0.11.0.0\\config\\server.properties"
	filename := "D:\\GoProject\\src\\goexamples\\tail\\my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,"+
				" filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		log.Println("msg :", msg.Text)
	}
}
