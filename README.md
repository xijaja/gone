<h1 align="center">Gone</h1>
<p align="center">远走高飞，前后端一次性部署，一个文件搞定</p>

## 特点

- 采用现代化技术栈
- 前后端开发时分离
- 单二进制程序构建
- 适合小型全栈项目
- 代码整洁注释齐全
- 单服务器轻量部署

## 介绍

本项目共有两个分支，`main` 分支和 `with-sqlite` 分支，两者在使用的数据库上有所不同。

两者均使用 embed 包，将前端打包后的文件嵌入到二进制文件中，最终部署时仅需部署该二进制文件。 `with-sqlite` 分支还将 sqlite 数据库嵌入。

早先，前端部分使用 Solid 和 DaisyUI，现在改用 SvelteKit 主要还是因为近期发布了 Svelte 5 这个带有符文（Runes）的版本。并且有神级 UI 库 ShadcnUI 加持，编写前端时会更加舒适。

大致技术栈：

- 前端：[Svelte/Kit](https://kit.svelte.dev) + TypeScript + [TailwindCSS](https://tailwindcss.com) + [ShadcnUI](https://github.com/shadcn-svelte/ui)
- 后端：Golang + [Fiber](https://github.com/gofiber/fiber) + [GORM](https://gorm.io) + Sqlite3/Postgres/MySQL + Dotenv + Air

最后，技术栈不是固定的，可以使用 Vue/React/Solid 等来替换 Svelte 前端框架，也可以使用 Rust 和 rust-embed 来替换 Golang。

## 功能模块

- 用户管理系统（登录等功能）
- Todo 待办事项管理（Demo）
- 静态文件服务（前端打包后的文件）
- 使用 GORM 作为 ORM 框架
- 支持多种数据库（Sqlite3/Postgres/MySQL）
- 使用 JWT 进行身份认证
- 包含日志中间件和 CSRF 保护
- 支持优雅关机
- 使用 embed 包将前端文件嵌入二进制文件

## 项目目录

项目目录也参考 [Go 项目标准布局](https://github.com/golang-standards/project-layout/blob/master/README_zh-CN.md) 进行了些许调整：

```sh
.
├── README.md                   # 项目介绍
├── Dockerfile                  # Dockerfile
├── deploy.sh                   # 部署脚本
├── makefile                    # makefile
├── go.mod                      # go.mod
├── go.sum                      # go.sum
├── LICENSE                     # 许可证
├── apis                        # api 路由
│   ├── handler                 # 路由处理
│   │   └── xxx.go              # xxx 路由处理
│   ├── middleware              # 中间件
│   │   ├── logs.go             # 日志中间件
│   │   └── xxx.go              # xxx 中间件
│   └── router.go               # 路由
├── cmd                         # 命令
│   ├── app                     # 程序入口
│   │   └── main.go             # 程序入口
│   └── cli                     # 命令行入口
│       └── main.go             # 命令行入口
├── config                      # 配置
│   ├── conf.go                 # 环境配置
│   └── start.go                # 启动配置
├── database                    # 数据库
│   ├── access                  # 数据库连接
│   │   ├── postgres.go         # pg 数据库连接
│   │   └── xxx.go              # xxx 数据库连接
│   └── model                   # 数据库模型
│       ├── auto_migrate.go     # 自动迁移
│       └── xxx_model.go        # xxx 模型
├── internal                    # 内部包
├── pkg                         # 外部包
└── svelte                      # 前端
    ├── frontend.go             # 嵌入前端
    ├── src                     # 前端源码
    │   ├── app.css             # app 样式
    │   ├── app.d.ts            # app 类型
    │   ├── app.html            # app html
    │   ├── lib                 # 库
    │   │   ├── components      # 组件
    │   │   ├── hooks           # 钩子
    │   │   ├── index.ts        # 入口
    │   │   └── utils.ts        # 工具
    │   └── routes              # 路由
    │       ├── +layout.svelte  # 布局
    │       ├── +layout.ts      # 布局
    │       └── +page.svelte    # 页面
    ├── static                  # 静态资源
    ├── package.json            # package.json
    ├── components.json         # ShadcnUI 组件配置
    ├── eslint.config.js        # ESLint 配置
    ├── postcss.config.js       # PostCSS 配置
    ├── svelte.config.js        # Svelte 配置
    ├── tailwind.config.ts      # TailwindCSS 配置
    ├── tsconfig.json           # TypeScript 配置
    └── vite.config.ts          # Vite 配置
```

## 食用方法

```sh
# 1.克隆项目
git clone https://github.com/xijaja/gone.git

# 2.进入项目
cd gone

# 3.复制 .env.dev 为 .env
cp .env.dev .env

# 4.编译前端（npm / yarn / pnpm）
cd svelte && pnpm install && pnpm run build

# 5.启动后端
cd .. && go run ./cmd/app

# 6.构建可执行文件（要记得编译前端）
go build -o app ./cmd/app

# 另外，构建不同的平台需要交叉编译
GOOS=linux GOARCH=amd64 go build -o app ./cmd/app        # linux amd64
GOOS=linux GOARCH=arm64 go build -o app ./cmd/app        # linux arm64
GOOS=darwin GOARCH=arm64 go build -o app ./cmd/app       # mac apple silicon
GOOS=darwin GOARCH=amd64 go build -o app ./cmd/app       # mac intel
GOOS=windows GOARCH=amd64 go build -o app.exe ./cmd/app  # windows amd64
GOOS=windows GOARCH=arm64 go build -o app.exe ./cmd/app  # windows arm64
```

## 部署相关

- 提供 deploy.sh 服务器部署脚本
- 提供 github actions 用于 fly.io 部署
- 提供 Dockerfile 容器化部署
- 提供 makefile 本地开发和容器命令

```sh
# 查看 makefile 中可用的命令
❯❯❯ make help
可用的命令:
  make pg               - 启动本地的 postgres 数据库
  make rs               - 启动本地的 redis 数据库
  make fly              - 部署到 fly.io
  make frontend         - 构建前端项目
  make build            - 构建 Docker 镜像
  make save             - 保存并传输镜像到远程主机
  make remote-run       - 在远程主机上运行容器
  make remote-ps        - 查看远程主机上的 Docker 状态
  make remote-images    - 查看远程主机上的镜像
  make deploy           - 执行完整的部署流程
  make clean            - 清理本地镜像
  make help             - 显示此帮助信息
```
