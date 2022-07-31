package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/event"
	"github.com/gorilla/websocket"
	"github.com/laizy/log"
	"github.com/r1cs/okx-api/common/utils"
	"sync/atomic"
)

type Client struct {
	Conn         *websocket.Conn
	readMsgFeed  event.Feed
	close        chan struct{}
	started uint32
}

type MsgEvent struct {
	Type int
	Data []byte
}

func (c *Client) readLoop() {
	for {
		select {
		case <-c.close:
			return

		default:
			mt, message, err := c.Conn.ReadMessage()
			if err != nil {
				log.Error("read message", "err", err)
				c.Close()
				return
			}
			c.readMsgFeed.Send(&MsgEvent{mt, message})
		}
	}
}

func (c *Client)SendMsg(msg *MsgEvent) {
		err:=c.Conn.WriteMessage(msg.Type,msg.Data)
		utils.Ensure(err)
}

func (c *Client) SubscribeReadMessage(ch chan<- *MsgEvent) event.Subscription {
	return c.readMsgFeed.Subscribe(ch)
}


var defalutDialer = websocket.DefaultDialer

func NewClient(url string) *Client {
	c, _, err := defalutDialer.Dial(url, nil)
	utils.Ensure(err)

	return &Client{
		Conn: c,
		close: make(chan struct{}),
	}
}

func(c *Client)Start()error{
	if !atomic.CompareAndSwapUint32(&c.started,0,1){
		return fmt.Errorf("already started")
	}
	go c.readLoop()
	return nil
}

func (c *Client) Close()error {
	close(c.close)
	return nil
}
