package main

import (
	"lan_api/conf"
	_ "lan_api/conf"
	"lan_api/server"
	"lan_api/util"
)

func main() {
	//fmt.Printf("%+v\n", conf.Data)
	util.Log().Info("main %+v\n", conf.Data)

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
