package net

import (
	"GameOfTCP/src/iface"
	"bytes"
	"encoding/binary"
)

type MsgPack struct{}

func NewMsgPack() *MsgPack {

	return &MsgPack{}
}
func (m *MsgPack) GetHeadLen() uint32 {
	//id和头长度各位4个字节
	return 8
}

//封包
func (m *MsgPack) MsgPack(msg iface.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, msg.GetDataLen())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

//拆包
func (m *MsgPack) MsgUnPack(byte []byte) (iface.IMessage, error) {

	reader := bytes.NewReader(byte)
	msg := &Message{}
	err := binary.Read(reader, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
