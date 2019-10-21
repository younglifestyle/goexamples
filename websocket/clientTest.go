package main

import (
	"fmt"
	"log"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Host string
	Path string
}

func NewWebsocketClient(host, path string) *Client {
	return &Client{
		Host: host,
		Path: path,
	}
}

var oneInt int32

func (this *Client) RecvMessage(query string) error {

	u := url.URL{Scheme: "ws", Host: this.Host, Path: this.Path, RawQuery: query}
	log.Printf("connecting to %s", u.String())

	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("1xx :", err)
		return err
	}
	defer ws.Close()

	for {
		//ws.Close()

		//err := ws.WriteControl(websocket.CloseMessage,
		//	websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		//	time.Now().Add(time.Second*10))
		//log.Println("1x one :", err)

		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("1x :", err)
			return err
		}

		fmt.Printf("%v %s received: %s\n", atomic.AddInt32(&oneInt, 1),
			query, message)
	}

	return nil
}

//172.16.9.229:18080
func main() {

	query := fmt.Sprintf("userid=%s&token=%s&id=%s",
		"1", "9a92f978a1e311e9af75005056b1076f", "13")

	//ws, _, err := websocket.DefaultDialer.Dial(uri, nil)
	//if err != nil {
	//	log.Println("1xx :", err)
	//	return
	//}
	//defer ws.Close()

	client := NewWebsocketClient("172.16.9.229:9967",
		"/api/v1/machine/ws")

	for i := 1; i <= 10; i++ {
		go client.RecvMessage(query)
		time.Sleep(50 * time.Millisecond)
	}

	select {}
}
