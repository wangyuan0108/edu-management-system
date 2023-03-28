#!/bin/bash

# 检测goimage(.gitlab-ci.yml提供的golang打包后的二进制文件的名称)是否已在运行, 如果已在运行则停止该容器
# docker ps 列出Docker容器列表
# grep goimage 搜索docker ps 列出Docker容器列表有没有goimage这个字串
# awk '{print $12}' 从docker ps 列出Docker容器列表的第12列(容器名)
# docker stop goimage 停止运行goimage容器
if [ $(docker ps | grep goimage | awk '{print $12}') ]; then docker stop goimage; fi

# 如果goimage容器存在 则删除该容器
# docker ps -a 列出全部容器
# docker ps -q 静默模式,只列出容器编号
# docker ps -aq 列出全部容器编号
# --filter name=goimage 过滤出容器名称为goimage的容器
# docker rm -f goimage 删除goimage容器
if [ $(docker ps -aq --filter name=goimage) ]; then
  docker rm -f goimage
fi

# docker images 列出所有镜像
# docker rmi -f 强制删除
# docker rmi -f goimage 强制删除goimage镜像
if [ $(docker images | grep goimage | awk '{print $1}') ]; then docker rmi -f goimage; fi

# 在Dockerfile当前目录的所有文件的镜像打包至标签为goimage的二进制文件可运行的Docker镜像
# docker build --tag 镜像的名字及标签，通常 name:tag 或者 name 格式；可以在一次构建中为一个镜像设置多个标签。

docker build --tag goimage -f /home/nginx/html/web/backend_file/Dockerfile .

# docker run 运行ddd
# -d 后台运行容器，并返回容器ID
# -p 映射主机与容器的端口
# --name 容器名称
docker run -d -p 4000:4000 --name goimage goimage

curl 192.168.0.158:4000/
curl 139.198.165.102:4000/
