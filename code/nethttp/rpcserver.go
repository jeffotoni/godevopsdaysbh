// Go in Action
// @jeffotoni

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var PORT_RPC = ":22334"

type Args struct {
	Json string
}

type Receive struct{}

func (t *Receive) Json(args *Args, reply *string) error {

	if len(args.Json) <= 0 {
		*reply = `{"status":"error", "msg":"json field is required"}`
		return nil
	}

	*reply = "ok"
	fmt.Println("publisher now: ", args.Json)
	return nil
}

func main() {
	re := new(Receive)
	serverRpc := rpc.NewServer()
	serverRpc.Register(re)
	serverRpc.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", PORT_RPC)
	if e != nil {
		log.Println("listen error:", e)
		return
	}

	for {
		if conn, err := listener.Accept(); err != nil {
			log.Println("accept error: ", err.Error())
			return
		} else {
			log.Printf("New connection established in rpc server\n")
			serverRpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}
