package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"testing"
)

// ---------------------------------------------
// 全局变量
// ---------------------------------------------
type projectConfig struct {
	FrontendStaticPath string   `validate:"required"` // 前端静态文件路径
	Postgres           Postgres // Postgres 数据库配置
	Redis              Redis    // Redis 配置
	JwtSecret          string   `validate:"required"` // JWT 密钥
	CsrfSecret         string   `validate:"required"` // CSRF 密钥
}

// Postgres 数据库配置
type Postgres struct {
	Host    string `validate:"required"` // IP地址
	User    string `validate:"required"` // 用户
	Port    string `validate:"required"` // 端口
	Pass    string `validate:"required"` // 密码
	Base    string `validate:"required"` // 库名
	Sslmode string `validate:"required"` // SSL模式
}

// Redis 配置
type Redis struct {
	Host string `validate:"required"` // IP地址
	Port string `validate:"required"` // 端口
	Pass string `validate:"required"` // 密码
	Base string `validate:"required"` // 库名
}

// 添加一个辅助函数来查找项目根目录
func findProjectRoot() string {
	// 从当前目录开始向上查找，直到找到 go.mod 文件所在的目录
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("获取当前工作目录失败: %v", err)
		return ""
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			// 已经到达根目录仍未找到
			return ""
		}
		dir = parent
	}
}

// 读取配置文件
func (mc *projectConfig) getMyConfig(isProd bool) projectConfig {
	if isProd {
		// 生产环境
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("生产环境加载 .env 文件时出错:", err.Error())
		}
	} else if !testing.Testing() {
		// 开发环境
		if err := godotenv.Load(".env.dev"); err != nil {
			log.Fatal("开发环境加载 .env.dev 文件时出错:", err.Error())
		}
	} else {
		// 测试环境
		rootDir := findProjectRoot()
		envPath := filepath.Join(rootDir, ".env.dev")
		if err := godotenv.Load(envPath); err != nil {
			log.Fatal("警告: 测试环境加载 .env.dev 文件失败:", err.Error())
		}
	}
	// 前端静态文件路径
	mc.FrontendStaticPath = os.Getenv("FRONTEND_STATIC_PATH")
	// Postgres 数据库配置
	mc.Postgres.Host = os.Getenv("PG_HOST")
	mc.Postgres.User = os.Getenv("PG_USER")
	mc.Postgres.Port = os.Getenv("PG_PORT")
	mc.Postgres.Pass = os.Getenv("PG_PASS")
	mc.Postgres.Base = os.Getenv("PG_BASE")
	mc.Postgres.Sslmode = os.Getenv("PG_SSLMODE")
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
