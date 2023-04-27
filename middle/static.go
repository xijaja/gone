package middle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"gone/frontend"
	"io/fs"
	"net/http"
)

func Pages(app *fiber.App) {
	// fs.Sub 用于获取嵌入文件系统的子目录
	dist, err := fs.Sub(frontend.FrontEnd, "dist")
	if err != nil {
		panic(err) // 嵌入文件系统有问题，地鼠们将会恐慌
	}

	// 使用 filesystem 中间件将 dist 目录作为静态文件目录
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(dist), // 使用 http.FS 包装 fs.FS
		Browse:       false,         // 不允许浏览目录
		Index:        "index.html",  // 默认访问 index.html
		MaxAge:       0,             // 3600 缓存 1 小时，单位秒，0 表示不缓存
		NotFoundFile: "index.html",  // 前端是 SPA 所以 404 时重定向到 index.html
	}))
}
