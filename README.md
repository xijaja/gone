# Gone

Gone 远走高飞，前后端一次性部署，一个文件搞定。

## 介绍

前端：Solid + TypeScript + TailwindCSS + DaisyUI

后端：Golang + Fiber + GORM + Sqlite3 + Dotenv + Air

使用 embed 包，将前端打包后的文件嵌入到二进制文件中（and Sqlite db file），实现一次性部署。

## 正确的食用方法

1. 当你拿到这个项目的代码时执行 `go run main.go` 肯定会报错，因为没有找到 /frontend/dist 文件夹。
   所以你需要先编译前端，比如执行 `cd frontend && npm run build`，然后再执行 `cd .. && go run main.go`。
   实际上，仅首次启动需要手动编译前端，在配置文件中天写你的包管理工具 npm、yarn 或 pnpm，
   之后的每次执行 `go run main.go` 我都会为你重新编译一次前端。
