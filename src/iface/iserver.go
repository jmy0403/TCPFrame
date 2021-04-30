package iface

//服务抽象层接口

type IServer interface {
	//准备启动服务器接口方法
	Start()

	//停止服务器接口方法
	Stop()

	//开启服务器接口方法
	Server()

	//
	AddRouter(msgId uint32, router IRouter)

	GetConnManager() IManageConn

	SetOnConnStart(func(IConn))
	//设置该Server的连接断开时的Hook函数
	SetOnConnStop(func(IConn))
	//调用连接OnConnStart Hook函数
	CallOnConnStart(conn IConn)
	//调用连接OnConnStop Hook函数
	CallOnConnStop(conn IConn)
}
