package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       666,
	}
	//连接Server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net Dial err:", err)
		return nil
	}

	client.conn = conn

	return client
}

func (this *Client) Menu() bool {
	var flag int
	fmt.Println("[1]公聊模式&&命令模式")
	fmt.Println("[2]私聊模式")
	fmt.Println("[3]更新用户名")
	fmt.Println("[0]退出系统")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println("请输入合法范围的数字")
		return false
	}
}

func (this *Client) Run() {
	for this.flag != 0 {
		//判断输入模式是否合法
		for this.Menu() != true {

		}
		switch this.flag {
		case 1:
			fmt.Println("进入公聊模式")
			this.PublicChat()
		case 2:
			fmt.Println("进入私聊模式")
			this.PrivateChat()
		case 3:
			fmt.Println("更新用户名")
			this.ReName()
		}
	}
}

func (this *Client) ReName() bool {
	fmt.Println("请输入更改后的新用户名:")
	fmt.Scanln(&this.Name)
	sendMsg := "rename " + this.Name + "\n"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return false
	}
	return true
}

func (this *Client) PublicChat() {
	var sendMsg string
	fmt.Println("请输入消息:(输入exit退出,输入help可以查看命令列表哦)")
	fmt.Scanln(&sendMsg)
	//TODO 添加空格适配
	for len(sendMsg) != 0 && sendMsg != "exit" {

		_, err := this.conn.Write([]byte(sendMsg + "\n"))
		if err != nil {
			fmt.Println("conn write err:", err)
			break
		}

		sendMsg = ""
		fmt.Println("请输入消息:(输入exit退出)")
		fmt.Scanln(&sendMsg)
	}
}

func (this *Client) PrivateChat() {
	var sendMsg string
	var objName string
	fmt.Println("请输入目标用户:(输入exit退出)")
	fmt.Scanln(&objName)
	if objName == "exit" || sendMsg == "exit" {
		return
	}
	fmt.Println("请输入消息:(输入exit退出)")
	fmt.Scanln(&sendMsg)
	if objName == "exit" || sendMsg == "exit" {
		return
	}
	//TODO 添加空格适配
	for objName != "exit" || sendMsg != "exit" {

		_, err := this.conn.Write([]byte("to " + objName + " " + sendMsg + "\n"))
		if err != nil {
			fmt.Println("conn write err:", err)
			break
		}

		sendMsg = ""
		fmt.Println("请输入目标用户:(输入exit退出)")
		fmt.Scanln(&objName)
		if objName == "exit" || sendMsg == "exit" {
			return
		}
		fmt.Println("请输入消息:(输入exit退出)")
		fmt.Scanln(&sendMsg)
		if objName == "exit" || sendMsg == "exit" {
			return
		}
	}
}

func (this *Client) DealResponse() {
	//阻塞检查conn的数据，并发送给标准输出
	io.Copy(os.Stdout, this.conn)
}
