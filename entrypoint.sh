#!/bin/bash

# 设置环境变量
export GIN_MODE=release
export PORT=8080

echo "当前工作目录: $(pwd)"
echo "环境变量:"
echo "GIN_MODE=$GIN_MODE"
echo "PORT=$PORT"

# 构建并运行应用
echo "开始构建应用..."
go build -o main .
echo "应用构建完成，开始运行..."
./main