# 使用官方的 Go 镜像作为基础镜像
FROM golang:1.21

# 安装 make 工具
RUN apt-get update && apt-get install -y --no-install-recommends \
    make \
    && rm -rf /var/lib/apt/lists/*

# 设置工作目录
WORKDIR /src

# 复制当前目录下的所有文件到容器的 /src 目录
COPY . /src

# 默认命令，这里使用 bash 进入交互模式
CMD ["bash"]
