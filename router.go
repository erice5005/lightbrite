package lightbrite

import (
	"log"
	"net"
	"strconv"
)

type Router struct {
	Conn        *net.UDPConn
	Port        int
	Addr        net.Addr
	Listen      *net.UDPConn
	WriteStream chan RouterFrame
	ReadStream  chan RouterFrame
}

type RouterFrame struct {
	Source  net.Addr
	Content interface{}
}

func NewRouter(targetIP net.IP, targetPort int) *Router {
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(targetIP.String(), strconv.Itoa(targetPort)))
	if err != nil {
		log.Print(err)
		return nil
	}
	log.Print(addr)
	// listAddr, _ := net.ResolveUDPAddr("udp", net.JoinHostPort("224.0.0.1", "0"))
	// log.Print(listAddr)
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Print(err)
		return nil
	}

	// l, err := conn.

	return &Router{
		Conn: conn,
		Port: targetPort,
		Addr: addr,
		// Listen:      l,
		WriteStream: make(chan RouterFrame, 10),
		ReadStream:  make(chan RouterFrame, 10),
	}
}

func (r *Router) Run() {
	go func() {
		for msg := range r.WriteStream {
			_, err := r.Conn.Write(msg.Content.(PilotRequest).toJSON())
			if err != nil {
				log.Printf("write err: %v\n", err)
			}
		}
	}()

	for {
		b := make([]byte, 2500)
		n, src, err := r.Conn.ReadFromUDP(b)
		if err != nil {
			continue
		}
		r.ReadStream <- RouterFrame{
			Source:  src,
			Content: PilotResponseFromJSON(b[:n]),
		}

	}
}
