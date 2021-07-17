#!/bin/bash

#编译文件
go build -o Server/Server Server/main.go Server/user.go Server/Server.go

echo "服务器启动..."

./Server/Server

