package main

import (
	"flag"
	"fmt"
)

var serverIp string
var serverPort int

func init() {
	//default:127.0.0.1
	flag.StringVar(&serverIp, "h", "127.0.0.1", "设置目标服务器IP(dafault:127.0.0.1)")
	//default:1234
	flag.IntVar(&serverPort, "p", 1234, "设置目标服务器Port(dafault:1234)")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("连接服务器失败")
		return
	}

	fmt.Println("服务器连接成功")

	go client.DealResponse()

	//启动客户端业务
	client.Run()
}
