package net

import (
	"GameOfTCP/src/iface"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

type Conn struct {
	TcpServer    iface.IServer
	Conn         *net.TCPConn
	ConnID       uint32
	IsClosed     bool
	Router       iface.IUserHandle
	ExitBuffChan chan bool
	msgChan      chan []byte
	msgBuffChan  chan []byte
	UserAttr     map[string]interface{}
	UserLock     sync.RWMutex
}

func (c *Conn) SetUserAttr(key string, value interface{}) {
	c.UserLock.Lock()
	defer c.UserLock.Unlock()
	c.UserAttr[key] = value
}

func (c *Conn) GetUserAttr(key string) interface{} {
	c.UserLock.RLock()
	defer c.UserLock.Unlock()
	return c.UserAttr[key]
}

func (c *Conn) RemoveUserAttr(key string) {
	c.UserLock.Lock()
	defer c.UserLock.Unlock()
	delete(c.UserAttr, key)

}

func (c *Conn) SendBuffMsg(msgId uint32, data []byte) error {
	if c.IsClosed == true {
		return errors.New("Connection closed when send buff msg")
	}
	//将data封包，并且发送
	dp := NewMsgPack()
	msg, err := dp.MsgPack(NewMessage(data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.msgBuffChan <- msg

	return nil
}

func NewConn(server iface.IServer, conn *net.TCPConn, id uint32, r iface.IUserHandle) *Conn {
	con := &Conn{
		TcpServer:    server,
		Conn:         conn,
		ConnID:       id,
		IsClosed:     false,
		Router:       r,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
		msgBuffChan:  make(chan []byte, 20),
		UserAttr:     make(map[string]interface{}),
	}
	con.TcpServer.GetConnManager().Add(con)
	return con

}

func (c *Conn) StartRead() {
	defer c.Stop()

	for {
		pack := NewMsgPack()
		dataHead := make([]byte, pack.GetHeadLen())
		_, err := io.ReadFull(c.Conn, dataHead)
		if err != nil {
			fmt.Println("从conn中读取数据失败", err)
			c.ExitBuffChan <- true
			continue

		}
		msg, err := pack.MsgUnPack(dataHead)
		if err != nil {
			fmt.Println("拆包失败", err)
			c.ExitBuffChan <- true
			continue
		}
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			_, err := io.ReadFull(c.Conn, data)
			if err != nil {
				fmt.Println("data拆包失败", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)
		req := Req{
			conn: c,
			data: msg,
		}

		c.Router.SendMsgToWorkQueue(&req)

	}
}

func (c *Conn) Start() {
	go c.StartRead()
	go c.StartWriter()
	c.TcpServer.CallOnConnStart(c)
	for {
		select {
		case <-c.ExitBuffChan:
			c.Stop()
			return
		}
	}
}
func (c *Conn) Stop() {
	if c.IsClosed == true {
		return
	}
	c.IsClosed = true
	c.TcpServer.CallOnConnStop(c)
	c.TcpServer.GetConnManager().CloseConn()
	c.Conn.Close()
	close(c.ExitBuffChan)
	close(c.msgChan)
}
func (c *Conn) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c *Conn) GetConnID() uint32 {
	return c.ConnID
}

func (c *Conn) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Conn) SendMsg(id uint32, data []byte) error {
	if c.IsClosed == true {
		return nil
	}

	pack := &MsgPack{}
	msg := NewMessage(data)
	msg.SetMsgId(id)
	msg.SetDataLen(uint32(len(data)))
	bytes, err := pack.MsgPack(msg)
	if err != nil {
		c.ExitBuffChan <- true
		return err
	}

	c.msgChan <- bytes
	return nil

}

func (c *Conn) StartWriter() {
	for {
		select {
		case msg := <-c.msgChan:
			_, err := c.Conn.Write(msg)
			if err != nil {
				fmt.Println("客户端写入数据失败", err)
			}
		case <-c.ExitBuffChan:
			return
		case data, ok := <-c.msgBuffChan:
			if ok {
				//有数据要写给客户端
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
					return
				}
			} else {
				break
				fmt.Println("msgBuffChan is Closed")
			}

		}
	}
}
