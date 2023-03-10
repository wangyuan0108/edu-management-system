# 从alpine镜像运行
FROM alpine

# 映射的端口
EXPOSE 4000
# 从宿主机./app复制到容器的目录/app
COPY . /

# 启动容器时运行app二进制文件
ENTRYPOINT ["/app"]
