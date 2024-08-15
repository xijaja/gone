#!/bin/zsh

# 读取环境变量
source .env.deploy

# 验证参数
# --------- --------- --------- --------- ---------
echo "🔍  验证参数"

# 判断多个环境变量是否存在
if [ -z "$SSH_USER" ] || [ -z "$SSH_IP" ] || [ -z "$SSH_KEY" ] || [ -z "$PROJECT_PATH" ] || [ -z "$PROJECT_NAME" ]; then
    echo "环境变量不存在，或部分不存在，请检查.env文件"
    exit 1
fi

# 如果 $FRONTEND_PATH 存在，则检查是否为目录
# shellcheck disable=SC2236
if [ ! -z "$FRONTEND_PATH" ] && [ ! -d "$FRONTEND_PATH" ] ; then
    echo "FRONTEND_PATH 必须为目录"
    exit 1
fi

# 如果 $FRONTEND_PATH 存在，那么 $FRONTEND_PKG_MANAGER 必须存在
# shellcheck disable=SC2236
if [ ! -z "$FRONTEND_PATH" ] && [ -z "$FRONTEND_PKG_MANAGER" ]; then
    echo "FRONTEND_PKG_MANAGER 必须存在"
    exit 1
fi

# SSH_IP 必须为IP地址
if [[ ! $SSH_IP =~ ^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "SSH_IP 必须为 IP 地址"
    exit 1
fi

# SSH_KEY 必须为文件路径，且文件存在
if [ ! -f "$SSH_KEY" ]; then
    echo "SSH_KEY 必须为文件路径，且文件存在"
    exit 1
fi

# PROJECT_PORT 必须为 0-65535 之间的数字
# shellcheck disable=SC2236
if [ ! -z "$PROJECT_PORT" ] && [[ ! $PROJECT_PORT =~ ^[0-9]+$ ]] && [[ $PROJECT_PORT -ge 0 ]] && [[ $PROJECT_PORT -le 65535 ]]; then
    echo "PROJECT_PORT 必须为 0-65535 之间的数字"
    exit 1
fi

# 开始部署
# --------- --------- --------- --------- ---------
echo "🚀  开始部署"

# 检查服务器上的程序是否存在，如果存在则删除
# shellcheck disable=SC2086
if ssh -i "$SSH_KEY" "$SSH_USER@$SSH_IP" -f $PROJECT_PATH/"$PROJECT_NAME"; then
    echo "🥷 服务器上存在程序，准备删除"
    # 删除服务器上的程序
    ssh -i "$SSH_KEY" "$SSH_USER@$SSH_IP" rm "$PROJECT_PATH"/"$PROJECT_NAME"
    echo "🗑️  删除完成"
else
    echo "服务器上不存在程序"
fi

# 检查本地的程序是否存在，如果存在则上传，如果不存在则编译项目然后再上传
if [ -f "$PROJECT_NAME" ]; then
    echo "⬆️ 本地存在程序，准备上传"
    # shellcheck disable=SC2140
    scp -i "$SSH_KEY" ./"$PROJECT_NAME" "$SSH_USER@$SSH_IP":"$PROJECT_PATH"/"$PROJECT_NAME"
else
    echo "🧬 本地不存在程序，准备编译项目"
    # 如果 $FRONTEND_PATH 存在，则编译前端程序
    if [ -d "$FRONTEND_PATH" ]; then
        echo "📦 编译前端程序..."
        cd "$FRONTEND_PATH" && $FRONTEND_PKG_MANAGER build && cd ..
    fi
    # 编译后端程序
    GOOS=linux GOARCH=amd64 go build -o "$PROJECT_NAME" -tags 'embed' main.go # 编译时嵌入静态文件
    # GOOS=linux GOARCH=amd64 go build -o $PROJECT_NAME main.go
    echo "✅  编译完成，正在上传中..."
    # shellcheck disable=SC2140
    scp -i "$SSH_KEY" ./"$PROJECT_NAME" "$SSH_USER@$SSH_IP":"$PROJECT_PATH"/"$PROJECT_NAME"
fi

# 重启服务器上的程序
echo "✅  重启服务器上的程序"

# 如果 $PROJECT_PORT 存在，则可以检查对应端口是否占用，然后杀死进程
# shellcheck disable=SC2236
if [ ! -z "$PROJECT_PORT" ]; then
    echo "🧐 检查对应端口是否占用"
    # 检查对应端口是否占用，如果占用则杀死进程，如果没有占用则不做任何操作
    ssh -i "$SSH_KEY" "$SSH_USER@$SSH_IP" nc -z 127.0.0.1 "$PROJECT_PORT" && ssh -i "$SSH_KEY" "$SSH_USER@$SSH_IP" "lsof -i:$PROJECT_PORT | awk 'NR==2 {print \$2}' | xargs kill"
else
    # 否则，直接杀死对应名称的进程
    ssh -i "$SSH_KEY" "$SSH_USER@$SSH_IP" "ps -ef | grep $PROJECT_NAME | grep -v grep | awk '{print \$2}' | xargs kill"
fi

# 重新运行服务器上的程序
# shellcheck disable=SC2087
ssh -i "$SSH_KEY" "$SSH_USER@$SSH_IP" >/dev/null 2>&1 <<eeooff
  cd $PROJECT_PATH
  nohup ./$PROJECT_NAME -s -t > log.txt 2>&1 &
  exit
eeooff


# 部署完成
# --------- --------- --------- --------- ---------
echo "🎉  部署完成"

# 清理本地程序
echo "✅  清理本地程序"
rm ./"$PROJECT_NAME"
echo "⛱️ 清理完成"
