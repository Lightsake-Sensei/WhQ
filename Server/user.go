package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

//创建一个用户对象
func NewUser(conn net.Conn, server *Server) *User {
	//获得当前连接的地址
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	//启动监听
	go user.ListenMessage()

	return user
}

//监听通道，有信息就发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

//用户上线功能
func (this *User) Online() {

	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

//用户下线功能
func (this *User) Offline() {

	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "下线")
}

//发送私有消息
func (this *User) PrivateSendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//发送新消息
func (this *User) DoMessage(msg string) {
	//私聊功能
	if len(msg) > 4 && msg[:3] == "to " {
		//获取用户名
		remoteName := strings.Split(msg, " ")[1]
		if remoteName == "" {
			this.PrivateSendMsg("请使用消息格式：to [Name] [Message]进行私聊\n")
			return
		}

		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.PrivateSendMsg("该用户名不存在\n")
			return
		}
		//发送消息内容给对方
		sendmsg := strings.Split(msg, " ")[2]
		if sendmsg == "" {
			this.PrivateSendMsg("请使用消息格式：to [Name] [Message]进行私聊\n")
			return
		}

		for i := 3; i < len(strings.Split(msg, " ")); i++ {
			sendmsg += " " + strings.Split(msg, " ")[i]
		}

		remoteUser.PrivateSendMsg("from[" + this.Name + "]:" + sendmsg + "\n")

		return
	}

	//重命名
	if len(msg) > 7 && msg[:7] == "rename " {
		newName := msg[7:]
		this.server.BroadCast(this, "rename:"+this.Name+"->"+msg[7:])
		//判断用户名是否存在
		if _, ok := this.server.OnlineMap[newName]; ok {
			this.PrivateSendMsg("当前用户已被使用\n")
		} else {
			this.server.mapLock.Lock()
			//换名
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this

			this.server.mapLock.Unlock()

			this.Name = newName
			this.PrivateSendMsg("用户名修改成功:" + this.Name + "\n")
		}
		return
	}
	//无附加属性命令集
	switch msg {
	case "who":
		//查询当前用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + "在线...\n"
			this.PrivateSendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	case "help":
		//查询指令
		this.PrivateSendMsg("---\nCommand: who \n\t Do: Search online user list\n")
		this.PrivateSendMsg("---\nCommand: rename [newName]\n\t Do: Modifications username\n")
		this.PrivateSendMsg("---\nCommand: to [Name] [Message]\n\t Do: Private chat to somebody\n")

	default:
		this.server.BroadCast(this, msg)
	}
}
