#!/bin/bash

# 需要创建的目录
dir=/home/nginx/html/web/temp
# 列出该目录所有文件
res=$(ls -A $dir)
# 检查该目录是否有文件, 没有则创建
if [ -z "$res" ]; then
  rm -rf /home/nginx/html/web/temp
else
  mkdir -p /home/nginx/html/web/temp
fi
