package start

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// ---------------------------------------------
// 全局变量
// ---------------------------------------------
type projectConfig struct {
	PkgManager string   `validate:"oneof='pnpm' 'cnpm' 'npm' 'yarn'"` // 前端的包管理器
	Database   Database // 数据库配置
}

// Database 数据库配置
type Database struct {
	Host string `validate:"required,ip"` // IP地址
	User string `validate:"required"`    // 用户
	Port string `validate:"required"`    // 端口
	Pass string `validate:"required"`    // 密码
	Base string `validate:"required"`    // 库名
}

// 读取配置文件
func (mc *projectConfig) getMyConfig(isProd bool) projectConfig {
	if isProd {
		if godotenv.Load(".env") != nil {
			log.Fatal("生产环境加载 .env 文件时出错")
		}
	} else {
		if godotenv.Load(".env.dev") != nil {
			log.Fatal("开发环境加载 .env.dev 文件时出错")
		}
	}
	// 前端包管理器
	mc.PkgManager = os.Getenv("PKG_MANAGER")
	// 读取配置信息
	mc.Database.Host = os.Getenv("DB_HOST")
	mc.Database.User = os.Getenv("DB_USER")
	mc.Database.Port = os.Getenv("DB_PORT")
	mc.Database.Pass = os.Getenv("DB_PASS")
	mc.Database.Base = os.Getenv("DB_BASE")
	// 验证配置是否正确
	validate := validator.New()
	if err := validate.Struct(mc); err != nil {
		log.Fatal(fmt.Sprintf("初始化验证配置: %v", err))
	}
	return *mc
}
