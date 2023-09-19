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
	PkgManager string `validate:"oneof='pnpm' 'cnpm' 'npm' 'yarn'"` // 前端的包管理器
	Postgres   struct {
		Host string `validate:"required,ip"` // IP地址
		User string `validate:"required"`    // 用户
		Port string `validate:"required"`    // 端口
		Pass string `validate:"required"`    // 密码
		Base string `validate:"required"`    // 库名
	} // 数据库配置
	Redis struct {
		Host string `validate:"required,ip"` // IP地址
		Port string `validate:"required"`    // 端口
		Pass string `validate:"required"`    // 密码
		Base string `validate:"required"`    // 库名
	} // Redis 配置
	JwtSecret  string `validate:"required"` // JWT 密钥
	CsrfSecret string `validate:"required"` // CSRF 密钥
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
	// Postgres 数据库配置
	mc.Postgres.Host = os.Getenv("PG_HOST")
	mc.Postgres.User = os.Getenv("PG_USER")
	mc.Postgres.Port = os.Getenv("PG_PORT")
	mc.Postgres.Pass = os.Getenv("PG_PASS")
	mc.Postgres.Base = os.Getenv("PG_BASE")
	// Redis 数据库配置
	mc.Redis.Host = os.Getenv("REDIS_HOST")
	mc.Redis.Port = os.Getenv("REDIS_PORT")
	mc.Redis.Pass = os.Getenv("REDIS_PASS")
	mc.Redis.Base = os.Getenv("REDIS_BASE")
	// 加密和密钥
	mc.JwtSecret = os.Getenv("JWT_SECRET")
	mc.CsrfSecret = os.Getenv("CSRF_SECRET")
	// 验证配置是否正确
	validate := validator.New()
	if err := validate.Struct(mc); err != nil {
		log.Fatal(fmt.Sprintf("初始化验证配置: %v", err))
	}
	return *mc
}
