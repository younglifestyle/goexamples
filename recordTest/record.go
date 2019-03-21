package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-resty/resty"
)

type Record struct {
	UID          string `json:"uid" binding:"required" gorm:"column:uid"`
	Serial       string `json:"serial" binding:"required" gorm:"column:serial"`
	Errcode      string `json:"errcode" binding:"required" gorm:"column:errcode"`
	BIN          string `json:"bin" binding:"required" gorm:"column:bin"`
	CreatedAt    int64  `json:"create_time" gorm:"column:time"`
	CurrentTime  int64  `json:"time,omitempty" gorm:"-"`
	Order        string `json:"order" binding:"required" gorm:"column:orders"`
	ControllerID string `json:"controller_id" binding:"required" gorm:"column:controller_id"`
	SocketID     string `json:"socket_id" binding:"required" gorm:"column:socket_id"`
	Flag         string `json:"flag" binding:"required" gorm:"column:flag"`
	Md5Key       string `json:"md5_key,omitempty" binding:"required" gorm:"column:md5_key"`
}

func main() {

	record := Record{
		UID:          "x",
		Serial:       "3rew3",
		Errcode:      "12fd3",
		BIN:          "214213",
		Order:        "TDRAM19010001",
		ControllerID: "13212",
		SocketID:     "2132e4",
		CurrentTime:  14921211397,
		Flag:         "xxxx",
		Md5Key:       "0",
	}

	rt := resty.New()

	for i := 4001; i <= 8000; i++ {

		rt.R().SetBody(record).
			Post("http://172.16.9.229:18080/api/v1/records/add")

		record.Md5Key = strconv.Itoa(i)
		record.UID = fmt.Sprintf("x%d", i)
		record.Serial = fmt.Sprintf("3rew3%d", i)
		record.BIN = fmt.Sprintf("214213%d", i)
		record.ControllerID = fmt.Sprintf("13212%d", i)
		record.SocketID = fmt.Sprintf("2132e4%d", i)
		record.Order = fmt.Sprintf("TDRAM1901000%d", rand.Int31n(2)+1)
		record.CurrentTime = int64(1552355799 + i)
		record.Flag = fmt.Sprintf("xxxx%d", i)

		fmt.Println(record)
		time.Sleep(time.Millisecond * 10)
	}
}
