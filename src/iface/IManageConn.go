package iface

type IManageConn interface {
	Add(conn IConn)
	Remove(conn IConn)
	Get(id uint32) (IConn, error)
	Length() int
	CloseConn()
}
