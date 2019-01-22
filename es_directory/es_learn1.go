package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

var (
	hostElastic = "172.16.9.229"
	portElastic = "9200"
)

// ElasticTypeRecord is the Elasticsearch type for ElasticRecords.
const ElasticTypeRecord = "record"

// ElasticIndex is the name of the Gutenfinder index.
const ElasticIndex = "produce"

// ;index
type Record struct {
	UID          string `json:"uid" binding:"required" gorm:"column:uid"`
	Serial       string `json:"serial" binding:"required" gorm:"column:serial"`
	Errcode      string `json:"errcode" binding:"required" gorm:"column:errcode"`
	BIN          string `json:"bin" binding:"required" gorm:"column:bin"`
	CreatedAt    int64  `json:"create_time" gorm:"column:time"`
	TimeStr      int64  `json:"time,omitempty" gorm:"-"`
	Order        string `json:"order" binding:"required" gorm:"column:orders"`
	ControllerID string `json:"controller_id" binding:"required" gorm:"column:controller_id"`
	SocketID     string `json:"socket_id" binding:"required" gorm:"column:socket_id"`
	Md5Key       string `json:"md5_key,omitempty" binding:"required" gorm:"column:md5_key"`
}

func arrStrToArrInterface(strSls []string) []interface{} {
	//[]interface{}(strSls)
	//cannot convert strSls (type []string) to type []interface {}

	newSls := make([]interface{}, len(strSls))
	for i, v := range strSls {
		newSls[i] = v
	}

	return newSls
}

func main() {

	ctx := context.TODO()

	// 连接ES
	client, err := elastic.NewClient(
		elastic.SetURL("http://" + hostElastic + ":" + portElastic))
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	// 插入一条新的记录
	//oneT := map[string]string{"ones": "bin1"}
	//resp, err := client.Index().Index(ElasticIndex).
	//	Type(ElasticTypeRecord).Id("1").BodyJson(oneT).Do(ctx)
	//if err != nil {
	//	fmt.Println("1 error :", err)
	//	return
	//}
	//fmt.Println(resp)

	//oneT := Record{UID: "12312", Serial: "1231"}
	//resp, err := client.Index().Index(ElasticIndex).
	//	Type(ElasticTypeRecord).BodyJson(oneT).Do(ctx)
	//if err != nil {
	//	fmt.Println("error :", err)
	//	return
	//}
	//fmt.Println(resp)

	//oneT := map[string]string{"bin": "bin1", "sn": "214325"}
	//strBytes, _ := json.Marshal(oneT)
	//fmt.Println("string :", string(strBytes))
	//resp, err := client.Index().Index(ElasticIndex).
	//	Type(ElasticTypeRecord).Id("123124").BodyString(string(strBytes)).Do(ctx)
	//if err != nil {
	//	fmt.Println("error :", err)
	//	return
	//}
	//fmt.Println(resp)

	// 插入多条记录
	//n := 0
	//for i := 0; i < 10; i++ {
	//	bulkRequest := client.Bulk()
	//	for j := 0; j < 10; j++ {
	//		n++
	//		//tweet := map[string]string{"user": "olivere",
	//		//	"message": "Package strconv implements conversions to " +
	//		//		"and from string representations of basic data types. " + strconv.Itoa(n)}
	//		tweet := Record{UID: "12312", Serial: "1231"}
	//		req := elastic.NewBulkIndexRequest().
	//			Index(ElasticIndex).Type(ElasticTypeRecord).
	//			Id(strconv.Itoa(n)).Doc(tweet)
	//		bulkRequest = bulkRequest.Add(req)
	//	}
	//	bulkResponse, err := bulkRequest.Do(ctx)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	if bulkResponse != nil {
	//
	//	}
	//	fmt.Println(i)
	//}

	//str := []string{"214325", "aaaaaa"}
	//arrInterface := arrStrToArrInterface(str)

	//aggs := elastic.NewTermsAggregation().Field("BIN.keyword").Size(10)
	//source, err := aggs.Source()
	//if err != nil {
	//	fmt.Println(err)
	//}

	termQuery := elastic.NewTermQuery("order_str.keyword", `aaaaaaaaaa`)
	//termQuery := elastic.NewTermQuery("field", "BIN.keyword")
	////termQuery1 := elastic.NewTermQuery("sn", "aaaaaa")
	searchResult, err := client.Search().
		Index(ElasticIndex). // search in index "twitter"
		Type(ElasticTypeRecord).
		Query(termQuery). /*Query(termQuery1).*/ // specify the query
		/*Sort("_id", true).*/ // sort by "user" field, ascending
		/*From(0). */ Size(0). // take documents 0-9
		Pretty(true).          // pretty print request and response JSON
		Do(ctx)                // execute
	if err != nil {
		fmt.Println("2 error :", err)
		return
	}

	fmt.Println("one list :", searchResult.Hits.TotalHits)

	// 查询插入的数据
	//termQuery := elastic.NewTermsQuery("sn", arrInterface...)
	//termQuery := elastic.NewTermQuery("field", "BIN.keyword")
	////termQuery1 := elastic.NewTermQuery("sn", "aaaaaa")
	//searchResult, err := client.Search().
	//	Index(ElasticIndex). // search in index "twitter"
	//	Type(ElasticTypeRecord).Aggregation("all_BIN", aggs).
	//	//Query(termQuery). /*Query(termQuery1).*/ // specify the query
	//	/*Sort("_id", true).*/ // sort by "user" field, ascending
	//	/*From(0). */ Size(0). // take documents 0-9
	//	Pretty(true).          // pretty print request and response JSON
	//	Do(ctx)                // execute
	//if err != nil {
	//	fmt.Println("2 error :", err)
	//	return
	//}
	//
	//bytes, _ := searchResult.Aggregations["all_BIN"].MarshalJSON()
	////fmt.Println("one :", string(bytes))
	//
	//var tmp RootParamStruct
	//err = json.Unmarshal(bytes, &tmp)
	//fmt.Println("112 one :", tmp)
	//
	//for _, value := range searchResult.Hits.Hits {
	//	fmt.Println("one :", string(*value.Source))
	//}
}

type SubParamStruct struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
}
type RootParamStruct struct {
	Buckets []SubParamStruct `json:"buckets"`
}
