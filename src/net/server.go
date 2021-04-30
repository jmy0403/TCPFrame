package net

import (
	"GameOfTCP/src/config"
	"GameOfTCP/src/iface"
	"errors"
	"fmt"
	"net"
	"time"
)

type Server struct {
	//服务器名称
	Name string
	//TCP版本
	IPVersion string
	//IP地址
	IP string
	//端口号
	Port        int
	handle      iface.IUserHandle
	Cm          iface.IManageConn
	OnConnStart func(conn iface.IConn)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn iface.IConn)
}

func (s *Server) SetOnConnStart(f func(iface.IConn)) {
	s.OnConnStart = f
}

func (s *Server) SetOnConnStop(f func(iface.IConn)) {
	s.OnConnStop = f
}

func (s *Server) CallOnConnStart(conn iface.IConn) {
	if s.OnConnStart != nil {
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn iface.IConn) {
	if s.OnConnStop != nil {
		s.OnConnStop(conn)
	}
}

func (s *Server) GetConnManager() iface.IManageConn {
	return s.Cm
}

func (s *Server) AddRouter(msgId uint32, router iface.IRouter) {

	s.handle.AddRouter(msgId, router)

}

//服务器句柄(服务器名称，TCP版本，IP地址，端口
func NewServer() iface.IServer {

	return &Server{
		Name:      config.Cfg.Name,
		IPVersion: "tcp4",
		IP:        config.Cfg.Host,
		Port:      config.Cfg.Port,
		handle:    NewUserHandle(),
		Cm:        NewManagerConn(),
	}
}

//开启网络服务
func (s *Server) Start() {
	fmt.Printf("[服务器: %s，IP地址: %s，端口号: %d]正在准备启动中",
		s.Name, s.IP, s.Port)

	go func() {
		s.handle.StartWorkPool()
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("服务器启动失败：", err)
			return
		}
		tcp, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("net.ListenTCP()调用失败", err)
			return
		}
		var cid uint32
		cid = 0
		for {
			conn, err := tcp.AcceptTCP()
			if err != nil {
				fmt.Println("tcp.AcceptTCP()调用失败", err)
				continue

			}
			if s.Cm.Length() >= config.Cfg.MaxConn {
				conn.Close()
				continue
			}
			newConn := NewConn(s, conn, cid, s.handle)
			cid++
			go newConn.Start()
		}

	}()

}

func Callback(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient ... ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

//停止网络服务
func (s *Server) Stop() {
	s.Cm.CloseConn()
}

//开始处理服务业务
func (s *Server) Server() {
	s.Start()
	for {
		time.Sleep(time.Second * 5)
	}

}
