package net

import "GameOfTCP/src/iface"

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(req iface.IReq)   {}
func (br *BaseRouter) Handle(req iface.IReq)      {}
func (br *BaseRouter) AfterHandle(req iface.IReq) {}
