package frontend

import "embed"

// FrontEnd 将 dist 文件夹作为前端静态文件目录
//
//go:embed dist/*
//go:embed dist/**/*
var FrontEnd embed.FS
