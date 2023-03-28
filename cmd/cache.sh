#!/bin/bash
# 该脚本用于超过一定次数删除缓存

# 运行一次Job tag自增一次
tag=$(cat ./tag.txt)
tag=$((tag+1))
echo $tag>./tag.txt

# 查看缓存列表
ls "$CI_PROJECT_DIR"/web/node_modules
ls "$CI_PROJECT_DIR"/web/client/node_modules
ls "$CI_PROJECT_DIR"/web/client/dist
ls "$CI_PROJECT_DIR"/app
ls "$CI_PROJECT_DIR"/vendor

# 如果缓存存在超过定义的次数则删除并重置tag为0次
#if [ "$TAG" -gt 0 ]; then
if(("$tag" == 0)); then
  echo "删除第7次产生的缓存"
  rm -rf "$CI_PROJECT_DIR"/web/node_modules
  rm -rf "$CI_PROJECT_DIR"/web/client/node_modules
  rm -rf "$CI_PROJECT_DIR"/web/client/dist
  rm -rf "$CI_PROJECT_DIR"/app
  rm -rf "$CI_PROJECT_DIR"/vendor
  tag=0
  echo $tag>./tag.txt
fi
