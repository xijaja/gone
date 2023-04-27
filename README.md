# Gone

Gone 远走高飞，前后端一次性部署，一个文件搞定。

## 介绍

前端：Solid + TypeScript + TailwindCSS + DaisyUI

后端：Golang + Fiber + GORM + Sqlite3 + Dotenv + Air

使用 embed 包，将前端打包后的文件嵌入到二进制文件中（and Sqlite db file），实现一次性部署。

## 项目目录

```
├── README.md                   # 项目介绍
├── main.go                     # 程序入口
├── go.mod                      # 项目依赖
├── Dockerfile                  # Dockerfile
├── docker-compose.yml          # docker-compose
├── apis                        # api 路由
│   ├── api.go                  # api 路由入口
│   └── xxx.go                  # xxx 路由
├── db                          # 数据库
│   ├── access                  # 数据库操作
│   │   ├── conn.go             # 数据库连接
│   │   └── sqlite.db           # sqlite 数据库文件
│   └── model                   # 数据库模型
│       └── xxx_model.go        # xxx 模型
├── frontend                    # 前端
│   ├── frontend.go             # 嵌入前端
│   ├── dist                    # 前端编译后的目录
│   ├── public                  # 静态资源
│   ├── src                     # 前端源码
│   │   ├── index.tsx           # 入口文件
│   │   ├── pages               # 页面
│   │   │   └── xxx.tsx         # xxx 页面
│   │   ├── auto-import.d.ts    # 自动导入
│   │   ├── router.tsx          # 前端路由
│   │   └── styles              # 样式目录
│   ├── index.html              # html 模板
│   ├── package.json            # package.json
│   ├── pnpm-lock.yaml          # pnpm-lock.yaml
│   ├── postcss.config.js       # postcss 配置
│   ├── tailwind.config.cjs     # tailwind 配置
│   ├── tsconfig.json           # ts 配置
│   └── vite.config.ts          # vite 配置
├── middle                      # 中间件
│   ├── logs.go                 # 日志中间件
│   └── static.go               # 静态文件中间件
├── start                       # 启动
│   ├── conf.go                 # 环境配置
│   └── start.go                # 启动配置
└── utils
```

## 食用方法

```shell
# 1.克隆项目
git clone https://github.com/xijaja/gone.git

# 2.进入项目
cd gone

# 3.编译前端（npm / yarn / pnpm）
cd frontend && pnpm install && pnpm run build

# 4.启动后端
cd .. && go run main.go

# 5.构建可执行文件
go build -o app -tags embed main.go

# 注意：构建不同的平台需要交叉编译
GOOS=linux GOARCH=amd64 go build -o app -tags embed main.go
GOOS=darwin GOARCH=arm64 go build -o app -tags embed main.go
GOOS=darwin GOARCH=amd64 go build -o app -tags embed main.go
GOOS=windows GOARCH=amd64 go build -o app.exe -tags embed main.go
```
