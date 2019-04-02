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

func (this *Client) RecvMessage(userid, token string) error {

	query := fmt.Sprintf("userid=%s&token=%s", userid, token)
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

		err := ws.WriteControl(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			time.Now().Add(time.Second*10))
		log.Println("1x one :", err)

		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("1x :", err)
			return err
		}

		fmt.Printf("%v %s received: %s\n", atomic.AddInt32(&oneInt, 1),
			userid, message)
	}

	return nil
}

//172.16.9.229:18080
func main() {
	client := NewWebsocketClient("172.16.9.229:18080", "/api/v1/notify/ws/message")

	for i := 1; i <= 1; i++ {
		go client.RecvMessage("8", "d96c061d4ae011e9bf7c005056b1076f")
		time.Sleep(50 * time.Millisecond)
	}

	select {}
}
