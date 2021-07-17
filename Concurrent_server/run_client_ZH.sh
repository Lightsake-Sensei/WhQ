#!/bin/bash

#编译文件
go build -o Client/Client Client/main.go Client/Client.go

echo "客户端启动..."

./Client/Client