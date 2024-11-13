package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"gone/config"
	"gone/svelte"
	"io/fs"
	"net/http"
)

// Pages 配置静态文件服务中间件
func Pages(app *fiber.App) {
	// fs.Sub 用于获取嵌入文件系统的子目录
	build, err := fs.Sub(svelte.Build, "build")
	if err != nil {
		panic(err) // 嵌入文件系统有问题，地鼠们将会恐慌
	}

	// 声明文件系统入口
	var FileRoot http.FileSystem
	if *config.S {
		// 生产环境：使用嵌入到二进制文件中的静态文件
		FileRoot = http.FS(build) // 生产环境使用 http.FS 包装 fs.FS
	} else {
		// 开发环境：直接使用本地文件系统中的静态文件
		FileRoot = http.Dir(config.Config.FrontendStaticPath) // 开发环境使用编译后的文件
	}

	// 使用 filesystem 中间件将 build 目录作为静态文件目录
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         FileRoot,     // 文件系统入口
		Browse:       false,        // 不允许浏览目录
		Index:        "index.html", // 默认访问 index.html
		MaxAge:       0,            // 3600 缓存 1 小时，单位秒，0 表示不缓存
		NotFoundFile: "index.html", // 前端是 SPA 所以 404 时重定向到 index.html
		// 添加错误处理
		// Next: func(c *fiber.Ctx) bool {
		// 	// 如果请求的是 API 路径，跳过静态文件处理
		// 	return strings.HasPrefix(c.Path(), "/api")
		// },
	}))
}
