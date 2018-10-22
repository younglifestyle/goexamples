package main

import (
	"encoding/json"
	"fmt"
)

var testData = []byte(`
{
	"index_no":"123",
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

func main() {
	var testMap map[string]interface{}

	json.Unmarshal(testData, &testMap)

	fmt.Println(testMap)
}
