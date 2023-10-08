#!/bin/zsh

# 脚本变量
user_server="user@123.345.456.678" # 服务器上的用户名和 IP 地址
project_path="/home/user/project-name" # 服务器上的项目路径
project_port="3030" # 服务器上的项目端口
# shellcheck disable=SC2088
key_path="~/.ssh/xx.pem" # 本地的密钥绝对路径

# 检查服务器上的程序是否存在，如果存在则删除
if ssh -i "${key_path}" "${user_server}" -f ${project_path}/app; then
  echo "服务器上存在 app 程序，准备删除"
  # 删除服务器上的程序
  ssh -i "${key_path}" "${user_server}" rm ${project_path}/app
  echo "🗑️  删除成功"
else
  echo "服务器上不存在 app 程序"
fi

# 检查本地的程序是否存在，如果存在则上传，如果不存在则编译项目然后再上传
if [ -f "./app" ]; then
  echo "本地存在 app 程序，准备上传"
  scp -i "${key_path}" ./app "${user_server}":${project_path}/app
else
  echo "本地不存在 app 程序，准备编译项目"
  echo "编译前端程序..."
  cd ./frontend && pnpm build && cd ..
  echo "编译后端程序..."
  GOOS=linux GOARCH=amd64 go build -o app -tags 'embed' main.go
  echo "✅  编译完成，正在上传中..."
  scp -i "${key_path}" ./app "${user_server}":${project_path}/app
fi

# 重启服务器上的 app 程序
echo "✅  重启服务器上的 app 程序"

# 检查对应端口是否占用，如果占用则杀死进程，如果没有占用则不做任何操作
ssh -i "${key_path}" "${user_server}" nc -z 127.0.0.1 ${project_port} && ssh -i "${key_path}" "${user_server}" "lsof -i:${project_port} | awk 'NR==2 {print \$2}' | xargs kill"

# 重新运行服务器上的 app 程序
# shellcheck disable=SC2087
ssh -i "${key_path}" "${user_server}" > /dev/null 2>&1 << eeooff
  cd ${project_path}
  nohup ./app -s > log.txt 2>&1 &
  exit
eeooff

# 清理本地 app 程序
echo "✅  清理本地 app 程序"
rm ./app
echo "✅  部署已全部完成"
