#web登录注册系统Dockerfile  用alpine:latest会小一点
FROM centos:latest

LABEL tanjl="tanjl@itdayang.com"

WORKDIR /app

ADD ./conf /app/conf
ADD ./Log /app/Log
ADD ./main /app/main
ADD ./static /app/static

EXPOSE 9090

ENTRYPOINT ["/app/main"]