#!/bin/bash

ls $CI_PROJECT_DIR/web/node_modules
ls $CI_PROJECT_DIR/web/client/node_modules
ls $CI_PROJECT_DIR/web/client/dist
ls $CI_PROJECT_DIR/app
ls $CI_PROJECT_DIR/vendor

#if [ "$TAG" -gt 0 ]; then
if(("$TAG" == 0)); then
  echo "删除第7次产生的缓存"
  rm -rf "$CI_PROJECT_DIR"/web/node_modules
  rm -rf "$CI_PROJECT_DIR"/web/client/node_modules
  rm -rf "$CI_PROJECT_DIR"/web/client/dist
  rm -rf "$CI_PROJECT_DIR"/app
  rm -rf "$CI_PROJECT_DIR"/vendor
fi
