package client

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/r1cs/okx-api/common/types"
	"github.com/r1cs/okx-api/common/utils"
	"testing"
)


func TestClient(t *testing.T){
	url :=  "wss://ws.okx.com:8443/ws/v5/public"
	c := NewClient(url)
	c.Start()
	ch := make(chan *MsgEvent,1)
	c.SubscribeReadMessage(ch)
	s := types.NewTradeSubscribeArgs("BTC-USDT")
	p := types.NewSubscribeTrade(s)
	b,err := json.Marshal(p)
	utils.Ensure(err)
	t.Log(string(b))
	c.SendMsg(&MsgEvent{Type: websocket.TextMessage,Data: b})
	for msg :=range ch{
	t.Log(string(msg.Data))
	}
	}
