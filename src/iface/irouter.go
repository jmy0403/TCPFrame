package iface

type IRouter interface {
	PreHandle(req IReq)
	Handle(req IReq)
	AfterHandle(req IReq)
}
