package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	//在线用户列表
	OnlineMap map[string]*User
	//互斥锁
	mapLock sync.RWMutex

	//消息广播channel
	Message chan string
}

//构造一个Server对象
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//启动服务器
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	//启动监听msg
	go this.ListenMessager()

	for {
		//accept
		//Accept以得到一个connect
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen accpet err:", err)
			continue
		}
		//do handler
		//用一个goroutine处理业务，每个connect分配一个goroutine
		go this.Handler(conn)
	}

	//close listen socket
	defer listener.Close()
}

func (this *Server) Handler(conn net.Conn) {
	//处理当前业务
	//用户上线并记录

	user := NewUser(conn, this)
	user.Online()

	//接受客户端消息
	go func() {
		//最多接受4096个字节
		buf := make([]byte, 4096)
		for {

			len, err := conn.Read(buf)
			if len == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			//正常得到数据
			//去除'\n'
			msg := string(buf[:len-1])

			//广播得到的数据
			user.DoMessage(msg)
		}
	}()

	//当前handle阻塞
	select {}
}

//广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

//监听广播信息并将其发送给全部在线User的goroutine
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		//将msg发送给全部在线的User
		this.mapLock.Lock()
		for _, client := range this.OnlineMap {
			client.C <- msg
		}
		this.mapLock.Unlock()
	}
}
