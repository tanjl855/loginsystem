#!/bin/bash

echo -e "****web登录注册系统，开始执行自动打包****\n\n"

git pull
echo -e "step1:合并代码完成\n\n"

export GO111MODULE=on
export GOPROXY="https://goproxy.cn"
export GOARCH=amd64
export GOOS=linux
echo -e "step2:环境变量配置完成\n\n"

go build main.go
echo -e "step3:编译完成\n\n"

chmod +x run.sh
chmod +x main
echo -e "step4:更改权限完成\n\n"

docker rm -f loginsystem
docker rmi loginsystem:latest
docker build -f /home/loginsystem/Dockerfile -t loginsystem .
echo -e "step5:docker构建镜像完成\n\n"


