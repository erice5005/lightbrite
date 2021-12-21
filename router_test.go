package lightbrite

import (
	"log"
	"net"
	"testing"
)

func Test_RouterConnect(t *testing.T) {
	rx := NewRouter(net.ParseIP("192.168.0.105"), 38899)
	if rx == nil {
		t.FailNow()
	}
	go rx.Run()

	initReq := PilotRequest{
		Method: "registration",
		Params: PilotParams{
			"register": true,
			"phoneMac": "AAAAAAAAAAAA",
			"phoneIp":  "1.2.3.4",
			"id":       "1",
		},
	}
	rx.WriteStream <- RouterFrame{
		Content: initReq,
	}

	// time.Sleep(1 * time.Minute)
}

func Test_GetPilot(t *testing.T) {
	rx := NewRouter(net.ParseIP("192.168.0.105"), 38899)
	if rx == nil {
		t.FailNow()
	}
	go rx.Run()

	initReq := PilotRequest{
		Method: "getPilot",
	}
	rx.WriteStream <- RouterFrame{
		Content: initReq,
	}
	select {
	case val := <-rx.ReadStream:
		log.Printf("Val: %v\n", val)
	}
}
