package main

import (
	"GameOfTCP/src/iface"
	"GameOfTCP/src/net"
	"fmt"
)

type Router struct {
}

func (r *Router) PreHandle(req iface.IReq) {

}

func (r *Router) AfterHandle(req iface.IReq) {

}

func (r *Router) Handle(req iface.IReq) {
	fmt.Println("Handle开始运行")
	err := req.GetConn().SendMsg(1, []byte("bag1"))
	if err != nil {
		fmt.Println(err)
	}
	err1 := req.GetConn().SendMsg(2, []byte("bag2"))
	if err != nil {
		fmt.Println(err1)
	}
	err2 := req.GetConn().SendMsg(3, []byte("bag3"))
	if err != nil {
		fmt.Println(err2)
	}

}
func main() {

	server := net.NewServer()
	server.AddRouter(&Router{})

	server.Server()
}
