package main

import (
	"GameOfTCP/src/iface"
	"GameOfTCP/src/net"
	"fmt"
)

type Router struct {
	net.BaseRouter
}

func (r *Router) Handle(req iface.IReq) {

	err := req.GetConn().SendBuffMsg(1, []byte("hello1"))
	if err != nil {
		fmt.Println(err)
	}

	err2 := req.GetConn().SendBuffMsg(2, []byte("hello2"))
	if err2 != nil {
		fmt.Println(err2)
	}

}
func DoStart(conn iface.IConn) {
	fmt.Println("在连接开始之前执行了")
}
func DoEnd(conn iface.IConn) {
	fmt.Println("连接断开后执行了")
}
func main() {
	server := net.NewServer()
	server.SetOnConnStart(DoStart)
	server.SetOnConnStop(DoEnd)
	server.AddRouter(1, &Router{})
	server.Server()
}
