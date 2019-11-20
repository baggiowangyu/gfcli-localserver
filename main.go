package main

import (
	_ "gf-app/boot"
	_ "gf-app/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	println("=========================gfcli-localserver=================================")
	println("本程序目的是在无互联网环境为gf-cli提供项目初始化服务，当前适应的gogf版本为v1.9.10")
	println("请将以下域名通过HOSTS文件映射到本地127.0.0.1:10086")
	println("1、goframe.org")
	println("2、gf.cdn.johng.cn")
	g.Server().Run()
}
