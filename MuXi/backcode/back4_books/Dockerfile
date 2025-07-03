# 使用官方的Go基础镜像
FROM golang:1.20-alpine

# 设置工作目录
WORKDIR /app

# 复制项目文件到工作目录
COPY . .

# 下载项目依赖
RUN go mod tidy

# 构建Go应用
RUN go build -o main .

# 暴露应用监听的端口
EXPOSE 8080

# 运行应用
CMD ["./main"]