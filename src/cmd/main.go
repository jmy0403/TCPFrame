package main

import (
	"GameOfTCP/src/iface"
	"GameOfTCP/src/net"
	"fmt"
)

type PingRouter struct {
	net.BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request iface.IReq) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConn().GetTCPConn().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	//创建一个server句柄
	s := net.NewServer()

	//配置路由
	s.AddRouter(&PingRouter{})

	//开启服务
	s.Server()
}
