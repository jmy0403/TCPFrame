package iface

import "net"

//TCP连接处理接口
type IConn interface {
	Start()
	Stop()
	GetTCPConn() *net.TCPConn
	GetConnID() uint32
	GetRemoteAddr() net.Addr
	SendMsg(id uint32, data []byte) error
	SendBuffMsg(msgId uint32, data []byte) error //添加带缓冲发送消息接口
	SetUserAttr(key string, value interface{})
	GetUserAttr(key string) interface{}
	RemoveUserAttr(key string)
}

//处理连接业务
type HandFunc func(*net.TCPConn, []byte, int) error
