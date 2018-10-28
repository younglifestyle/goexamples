package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

var testDas = []byte(`
{
	"index_no": 1,
	"report_info": {
		"department":"development",
		"reporter":"one",
		"time":"2018-10-17 10:33:00",
		"test1":"1",
		"test3":"1",
		"test2":"1",
		"test212":{
			"one":"123",
			"arew":"321"
		}
	},
	"report_describe": {
		"department":"development",
		"reporter":"one",
		"time":"2018-10-17 10:33:00"
	},
	"test_result": {
		"department":"development",
		"reporter":"one",
		"time":"2018-10-17 10:33:00"
	}
}`)

type testImportStruct struct {
	IndexNo        int         `json:"index_no"`
	ReportInfo     interface{} `json:"report_info"`
	ReportDescribe interface{} `json:"report_describe"`
	TestResult     interface{} `json:"test_result"`
}

func main() {
	var tsStruct testImportStruct
	err := json.Unmarshal(testDas, &tsStruct)
	if err != nil {
		log.Println("json is error")
		return
	}

	openFile, _ := os.OpenFile("./datastore.json",
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	i := 0
	for i < 1000000 {
		tsStruct.IndexNo += i
		bytes, _ := json.Marshal(tsStruct)

		openFile.WriteString(string(bytes) + "\n")

		i++
		time.Sleep(time.Microsecond * 100)
	}

	openFile.Close()
}
