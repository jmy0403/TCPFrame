package iface

type IReq interface {
	GetConn() IConn
	GetData() []byte
	GetMsgId() uint32
}
