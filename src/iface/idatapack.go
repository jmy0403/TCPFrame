package iface

type IDataPack interface {
	GetHeadLen() uint32
	MsgPack(msg IMessage) ([]byte, error)
	MsgUnPack([]byte) (IMessage, error)
}
