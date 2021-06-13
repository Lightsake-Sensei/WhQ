package main

import "flag"

var serverIp string
var serverPort int

func init() {
	//default:127.0.0.1
	flag.StringVar(&serverIp, "h", "127.0.0.1", "设置当前服务器IP(dafault:127.0.0.1)")
	//default:1234
	flag.IntVar(&serverPort, "p", 1234, "设置当前服务器Port(dafault:1234)")
}

func main() {
	flag.Parse()

	server := NewServer(serverIp, serverPort)
	server.Start()
}
