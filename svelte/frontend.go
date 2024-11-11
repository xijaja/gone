package svelte

import "embed"

// Build 将 build 文件夹作为前端静态文件目录
//
//go:embed build/*
//go:embed build/**/*
var Build embed.FS
