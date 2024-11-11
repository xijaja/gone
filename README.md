<h1 align="center">Gone</h1>
<p align="center">远走高飞，前后端一次性部署，一个文件搞定</p>

## 特点

- 前后端开发时分离
- 单二进制程序构建
- 适合小型全栈项目
- 单服务器轻量部署

## 介绍

本项目共有两个分支，`main` 分支和 `with-sqlite` 分支，两者在使用的数据库上有所不同。

两者均使用 embed 包，将前端打包后的文件嵌入到二进制文件中，最终部署时仅需部署该二进制文件。 `with-sqlite` 分支还将 sqlite 数据库嵌入。

早先，前端部分使用 Solid 和 DaisyUI，现在改用 SvelteKit 主要还是因为近期发布了 Svelte 5 这个带有符文（Runes）的版本。并且有神级 UI 库 ShadcnUI 加持，编写前端时会更加舒适。

大致技术栈：

- 前端：Svelte/Kit + TypeScript + TailwindCSS + ShadcnUI
- 后端：Golang + Fiber + GORM + Sqlite3/Postgres/MySQL + Dotenv + Air

最后，技术栈不是固定的，可以使用 Vue/React/Solid 等来替换 Svelte 前端框架，也可以使用 Rust 和 rust-embed 来替换 Golang。

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

# 3.编译前端（npm / yarn / pnpm）
cd svelte && pnpm install && pnpm run build

# 4.启动后端
cd .. && go run ./cmd/app

# 5.构建可执行文件（要记得编译前端）
go build -o app ./cmd/app

# 另外，构建不同的平台需要交叉编译
GOOS=linux GOARCH=amd64 go build -o app ./cmd/app        # linux amd64
GOOS=linux GOARCH=arm64 go build -o app ./cmd/app        # linux arm64
GOOS=darwin GOARCH=arm64 go build -o app ./cmd/app       # mac apple silicon
GOOS=darwin GOARCH=amd64 go build -o app ./cmd/app       # mac intel
GOOS=windows GOARCH=amd64 go build -o app.exe ./cmd/app  # windows amd64
GOOS=windows GOARCH=arm64 go build -o app.exe ./cmd/app  # windows arm64
```
