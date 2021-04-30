package net

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(u uint32) {
	m.Id = u
}

func (m *Message) SetData(bytes []byte) {
	m.Data = bytes
}

func (m *Message) SetDataLen(u uint32) {
	m.DataLen = u
}
func NewMessage(data []byte) *Message {
	var id uint32
	id = 00001
	m := &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
	id++
	return m
}
