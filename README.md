# Gone

Gone 远走高飞，前后端一次性部署，一个文件搞定。

## 介绍

本项目共有两个分支，`main` 分支和 `with-sqlite` 分支，两者仅在使用的数据库上有所不同。

两者均使用 embed 包，将前端打包后的文件嵌入到二进制文件中，最终部署时仅需部署该二进制文件。

如果使用 Postgres/MySQL/Mongo/Redis 等数据库在服务器上使用 docker 启动即可。

大致技术栈：

- 前端：Solid + TypeScript + TailwindCSS + DaisyUI
- 后端：Golang + Fiber + GORM + Sqlite3/Postgres/MySQL + Dotenv + Air

最后，技术栈不是固定的，可以使用 vue/react/svelte 等来替换 solid 前端框架，也可以使用 rust 和 rust-embed 来替换 golang。

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

```
# 1.克隆项目

git clone https://github.com/xijaja/gone.git

# 2.进入项目

cd gone

# 3.编译前端（npm / yarn / pnpm）

cd frontend && pnpm install && pnpm run build

# 4.启动后端

cd .. && go run main.go

# 5.构建可执行文件

go build -o app main.go

# 不过，要记得编译前端

cd frontend && pnpm build && cd .. && go build -o app main.go

# 另外，构建不同的平台需要交叉编译

GOOS=linux GOARCH=amd64 go build -o app main.go        # linux
GOOS=darwin GOARCH=arm64 go build -o app main.go       # mac m1
GOOS=darwin GOARCH=amd64 go build -o app main.go       # mac intel
GOOS=windows GOARCH=amd64 go build -o app.exe main.go  # windows
```