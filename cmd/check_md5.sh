#!/bin/bash

md5sum -c --status d.md5

# 获取上一条命令的返回值
rc=$?

# 判断返回值并输出相应信息
if [ $rc -eq 0 ]; then
  echo "所有文件MD5值校验成功"
else
  echo "文件校验失败!"
  md5sum -c d.md5 | grep -v "OK"
  exit
fi
