# 第一阶段：构建阶段
# 使用 golang:alpine 作为基础镜像,并命名为 builder
FROM golang:alpine AS builder

# 设置工作目录
WORKDIR /go/src/building

# 将项目目录下的所有文件复制到容器的工作目录
COPY . .

# 使用 goproxy 代理（可选）
# RUN go env -w GOPROXY="https://goproxy.cn,direct"

# 下载项目依赖
RUN go mod download

# 构建 Go 应用程序
# 编译应用 -ldflags="-s -w" 用于减小二进制文件大小
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app ./cmd/app/main.go

# 第二阶段：运行阶段
# 使用 alpine 作为基础镜像以减小最终镜像大小
FROM alpine:latest

# 设置时区为北京时间
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

# 从 builder 阶段复制编译好的二进制文件
COPY --from=builder /go/src/building/app ./app

# 从 builder 阶段复制 .env 文件
COPY --from=builder /go/src/building/.env .

# 声明容器将监听的端口
EXPOSE 3030

# 设置入口点和默认命令
ENTRYPOINT ["/app"]
CMD ["-s"]