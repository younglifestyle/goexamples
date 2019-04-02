package main

import (
	"log"
	"strconv"
	"time"
)

// Client is a middleman between the websocket connection and the hub.
type WsClient struct {

	// userid flag
	Userid int

	Send chan []byte
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 5 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 1) / 10
)

func (ws *WsClient) WriteMsg() {
	//var unRegister bool
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		log.Println("write message ：", ws.Userid)
		ticker.Stop()
	}()

	for {
		select {
		case message, _ := <-ws.Send:
			log.Println("client get message :", string(message))

		case <-ticker.C:
			log.Println("ticker run...")
			//return
		}
	}
}

func main() {

	wsMap := make(map[int]*WsClient)

	for i := 1; i <= 5; i++ {
		wsMap[i] = &WsClient{
			Userid: 1,
			Send:   make(chan []byte, 10),
		}
		go wsMap[i].WriteMsg()
	}

	for i := 1; i <= 5; i++ {
		wsMap[i].Send <- []byte("test" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		close(wsMap[i].Send)
	}

	select {}
}
