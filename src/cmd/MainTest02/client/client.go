package main

import (
	"GameOfTCP/src/net"
	"fmt"
	"io"
	tcp "net"
	"time"
)

func main() {
	fmt.Println("client正在准备运行")
	conn, err := tcp.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("客户端连接服务器失败", err)
	}
	for {
		dp := net.NewMsgPack()
		message := net.NewMessage([]byte("aa"))
		message.Id = 1
		pack, _ := dp.MsgPack(message)
		_, err := conn.Write(pack)
		if err != nil {
			fmt.Println(err)
		}
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error", err)
			break
		}
		//将headData字节流 拆包到msg中
		msgHead, err := dp.MsgUnPack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*net.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			fmt.Println(" Msg: ID=", msg.Id, ", len:", msg.DataLen, ", data:", string(msg.Data))
		}

		time.Sleep(1 * time.Second)
	}
}
