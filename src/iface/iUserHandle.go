package iface

type IUserHandle interface {
	DoUserHandle(req IReq)
	AddRouter(msgId uint32, r IRouter)
	StartWorkPool()
	SendMsgToWorkQueue(req IReq)
}
