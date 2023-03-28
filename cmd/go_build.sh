#!/bin/bash
# 该脚本用于Go项目部署

echo "传输Golang文件除了web前端目录,.cache缓存,.ideaIDE缓存,tmp缓存目录以外的所有文件发送至主机root@139.198.165.102的/home/nginx/html/web/backend_file目录下"
sshpass -e rsync -a -e "ssh -o stricthostkeychecking=no" --exclude={web,.git,.cache,.idea,tmp} ./ root@139.198.165.102:/home/nginx/html/web/backend_file
sshpass -e rsync -a -e "ssh -o stricthostkeychecking=no" --exclude={web,.git,.cache,.idea,tmp} ./ root@192.168.0.158:/home/nginx/html/web/backend_file

echo "进入到主机root@139.198.165.102的/home/nginx/html/web/backend_file下执行./cmd/go_build.sh Golang部署script,执行完成之后删除/home/nginx/html/web/backend_file缓存目录"
sshpass -e ssh -o stricthostkeychecking=no root@139.198.165.102 'cd /home/nginx/html/web/backend_file && bash ./cmd/go_build.sh && rm -rf /home/nginx/html/web/backend_file && rm -rf /home/nginx/html/temp/dir.sh && rm -rf /home/nginx/html/web/backend_file'
sshpass -e ssh -o stricthostkeychecking=no root@192.168.0.158 'cd /home/nginx/html/web/backend_file && bash ./cmd/go_build.sh && rm -rf /home/nginx/html/web/backend_file && rm -rf /home/nginx/html/temp/dir.sh && rm -rf /home/nginx/html/web/backend_file'
