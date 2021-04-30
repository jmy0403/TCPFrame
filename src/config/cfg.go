package config

import (
	"GameOfTCP/src/iface"
	"encoding/json"
	"io/ioutil"
)

type GlobalCfg struct {
	TcpServer     iface.IServer //当前全局Server对象
	Host          string        //当前服务器主机IP
	Port          int           //当前服务器主机监听端口号
	Name          string        //当前服务器名称
	Version       string        //当前版本号
	MaxPacketSize uint32        //都需数据包的最大值
	MaxConn       int           //当前服务器主机允许的最大链接个数
	PoolSize      uint32
}

var Cfg *GlobalCfg

func (cfg *GlobalCfg) Reload() {

	file, err := ioutil.ReadFile(".\\src\\config\\cfg.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, Cfg)
	if err != nil {
		panic(err)
	}
}
func init() {
	Cfg = &GlobalCfg{

		Host:          "127.0.0.1",
		Port:          8080,
		Name:          "",
		Version:       "v1.0.0",
		MaxPacketSize: 10086,
		MaxConn:       5000,
		PoolSize:      50,
	}
	Cfg.Reload()
}
