# WhQ

---

a simple chat tool with thread

---

## 如何使用
`./run_server.sh` 启动服务器(默认设置127.0.0.1:1234)
`./run_client.sh` 启动客户端(默认设置127.0.0.1:1234)

**或使用**
`Server/Server -h`
`Client/Client -h` 查看帮助

`Server/Server -h [ip] -p [port]` 启动服务器
`Client/Client -h [ip] -p [port]` 启动客户端

**或直接用nc连接**

---
## 存在的问题

**目前使用client连接无法使用空格,使用nc可以使用空格**

`//TODO 防止用户输入信息因为接收信息而分行`

