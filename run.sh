#!/bin/bash


docker rm -f loginsystem
docker run -itd --name loginsystem -v /home/gowebdata:/app -p 9090:9090 loginsystem
echo -e "step1:loginsystem执行完成\n\n"