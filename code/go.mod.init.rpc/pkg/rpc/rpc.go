package rpcx

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var PORT_RPC = ":22334"

type Args struct{ Json string }
type Receive struct{}

func (t *Receive) Json(args *Args, reply *string) error {
	if len(args.Json) <= 0 {
		*reply = `{"status":"error", "msg":"json field is required"}`
		return nil
	}

	*reply = `{"status":"ok", "msg":"Receive json"}`
	log.Println(args.Json)
	return nil
}

func Run() {
	re := new(Receive)
	serverRpc := rpc.NewServer()
	serverRpc.Register(re)
	serverRpc.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", PORT_RPC)
	if e != nil {
		log.Println("listen error:", e)
		return
	}

	// to listen
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Println("accept error: " + err.Error())
			return
		} else {
			serverRpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}
