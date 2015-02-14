package plugin

import (
	"net"
	"net/rpc"
	"log"
)

type Args struct {
	X, Y int
}

type Science struct {
}

func (t *Science) Add(args *Args, reply *int) error {
	*reply = (args.X+args.Y)*2
	return nil
}

type Calculator struct{
	Science *Science
}

func (t *Calculator) Add(args *Args, reply *int) error {
	*reply = args.X+args.Y
	return nil
}

func init() {
	cal := &Calculator{&Science{}}

	rpc.RegisterName("aaa", cal)
	rpc.RegisterName("bbb", cal.Science)

	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	go func() {
		for {
			if conn, err := listener.Accept(); err != nil {
				log.Fatal("accept error: " + err.Error())
			} else {
				log.Printf("new connection established\n")
				go rpc.ServeConn(conn)
			}
		}
	}()
}

