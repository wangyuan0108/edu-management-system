#!/bin/bash

# 需要创建的目录
dir=/home/nginx/html/web/temp
# 列出该目录所有文件
res=$(ls -A $dir)
# 检查该目录是否有文件, 没有则创建
if [ -z "$res" ]; then
  echo "已存在临时目录,正在删除"
  rm -rf /home/nginx/html/web/temp
fi
echo "正在创建临时目录"
  mkdir -p /home/nginx/html/web/temp
