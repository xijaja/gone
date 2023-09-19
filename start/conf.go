package start

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// ---------------------------------------------
// 全局变量
// ---------------------------------------------
type projectConfig struct {
	PkgManager string `validate:"oneof='pnpm' 'cnpm' 'npm' 'yarn'"` // 前端的包管理器
}

// 读取配置文件
func (mc *projectConfig) getMyConfig(isProd bool) projectConfig {
	loadFileName := ".env.dev" // 默认为开发环境文件
	if isProd {
		loadFileName = ".env" // 生产环境时更换为生产环境文件
	}
	if godotenv.Load(loadFileName) != nil {
		// log.Fatal("加载环境变量文件时出错：", loadFileName)
		mc.PkgManager = "npm" // 默认使用 npm 作为前端包管理器
	} else {
		mc.PkgManager = os.Getenv("PKG_MANAGER") // 从文件中获取前端包管理器
	}
	// 验证配置是否正确
	validate := validator.New()
	if err := validate.Struct(mc); err != nil {
		log.Panicln("初始化验证配置: ", err)
	}
	return *mc
}
