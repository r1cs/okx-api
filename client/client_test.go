package client

import "testing"

func TestClient(t *testing.T){
	url :=  "wss://ws.okx.com:8443/ws/v5/public"
	c := NewClient(url)
	c.Start()
	ch := make(chan *MsgEvent,1)
	c.SubscribeReadMessage(ch)
	for msg :=range ch{
		t.Log(msg)

	}
	}
