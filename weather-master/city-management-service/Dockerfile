# 使用Go官方镜像
FROM golang:1.19 as builder

# 设置环境变量
ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn,direct \
CGO_ENABLED=0

# 设置工作目录
WORKDIR /app

# 复制项目文件到容器中
COPY . .

# 构建应用程序
RUN go build -o main .

# 暴露端口
EXPOSE 8001

# 设置容器启动时执行的命令
CMD ["./main"]
