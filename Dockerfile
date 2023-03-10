# 使用 docker build 构建镜像
FROM golang:alpine AS builder

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR $GOPATH/src/building

# 将当前目录下的所有文件复制到工作目录中
COPY . .

# 下载依赖
RUN go mod tidy

# 编译
RUN go build -o oneapp .

# 使用最小的镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 将编译好的二进制文件复制到镜像中
# 如果有配置文件，应在此时复制到镜像中
COPY --from=builder $GOPATH/src/building/oneapp .

# 暴露端口
EXPOSE 3000

# 暴露目录
VOLUME /app/logs

# 运行 app
CMD ["nohup /app/oneapp >> /app/logs/oneapp.log 2>&1 &"]


# 此 Dockerfile 结合 docker-compose.yml 一同使用，但也可以单独使用
# 单独构建镜像：
# 构建推送之前应先在 Docker Hub 上注册该同名镜像
# 构建镜像：当前平台构建
# docker build -t oneapp:0.1 .
# 构建镜像：多平台构建，“--push” 推送到 hub，另外 “-o type=registry” 是 type=image,push=true 的精简表示
# docker buildx build -t xijaja/oneapp:0.1 --platform=linux/arm64,linux/amd64,windows/amd64 . --push
# 拉取镜像：
# docker pull xijaja/oneapp:0.1
# 启动容器（直接执行该命令时，若本地无该镜像则将自动拉取）：
# docker run -itd -p 5000:5000 --name resume xijaja/oneapp:0.1
