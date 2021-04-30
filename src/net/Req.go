package net

import "GameOfTCP/src/iface"

type Req struct {
	conn iface.IConn
	data iface.IMessage
}

func (r *Req) GetConn() iface.IConn {
	return r.conn
}
func (r *Req) GetData() []byte {
	return r.data.GetData()
}

func (r *Req) GetMsgId() uint32 {
	return r.data.GetMsgId()
}
