package net

import (
	"GameOfTCP/src/config"
	"GameOfTCP/src/iface"
	"fmt"
	"strconv"
)

type UserHandle struct {
	Apis      map[uint32]iface.IRouter
	PoolSize  uint32
	WorkQueue []chan iface.IReq
}

func NewUserHandle() *UserHandle {
	return &UserHandle{
		Apis:      make(map[uint32]iface.IRouter),
		PoolSize:  config.Cfg.PoolSize,
		WorkQueue: make([]chan iface.IReq, config.Cfg.PoolSize),
	}
}

func (u *UserHandle) DoUserHandle(req iface.IReq) {
	id := req.GetMsgId()
	r, ok := u.Apis[id]
	if !ok {
		fmt.Println("找不到用户对应的方法")
		return
	}
	r.PreHandle(req)
	r.AfterHandle(req)
	r.Handle(req)
}

func (u *UserHandle) AddRouter(msgId uint32, r iface.IRouter) {
	if _, ok := u.Apis[msgId]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}
	//2 添加msg与api的绑定关系
	u.Apis[msgId] = r

}

func (u *UserHandle) StartOneWorkPool(id int, taskQueue chan iface.IReq) {
	for {

		select {
		case req := <-taskQueue:
			u.DoUserHandle(req)

		}
	}

}
func (u *UserHandle) StartWorkPool() {
	for i := 0; i < int(u.PoolSize); i++ {
		u.WorkQueue[i] = make(chan iface.IReq, config.Cfg.PoolSize)
		go u.StartOneWorkPool(i, u.WorkQueue[i])
	}
}
func (u *UserHandle) SendMsgToWorkQueue(request iface.IReq) {
	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则

	//得到需要处理此条连接的workerID
	workerID := request.GetConn().GetConnID() % u.PoolSize
	fmt.Println("Add ConnID=", request.GetConn().GetConnID(), " request msgID=", request.GetMsgId(), "to workerID=", workerID)
	//将请求消息发送给任务队列
	u.WorkQueue[workerID] <- request
}
