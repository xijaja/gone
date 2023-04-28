package start

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// ---------------------------------------------
// 全局变量
// ---------------------------------------------
type projectConfig struct {
	PkgManager string // 前端的包管理器
	NeedBuild  bool   // 是否需要构建前端
	UsePgSQL   bool   // 是否使用 PgSQL
	PgSQL      struct {
		Addr string // IP地址
		User string // 用户
		Port string // 端口
		Pass string // 密码
		Base string // 库名
	} // pgsql 配置
}

// 读取配置文件
func (mc *projectConfig) getMyConfig(isProd bool) projectConfig {
	if isProd {
		if godotenv.Load(".env") != nil {
			log.Fatal("生产环境加载 .env 文件时出错")
		}
		// 前端包管理器
		mc.PkgManager = os.Getenv("PKG_MANAGER")
		mc.NeedBuild = os.Getenv("NEED_BUILD") == "true"
		// PgSQL
		mc.UsePgSQL = os.Getenv("USE_PGSQL") == "true"
		mc.PgSQL.Addr = os.Getenv("PGSQL_ADDR")
		mc.PgSQL.User = os.Getenv("PGSQL_USER")
		mc.PgSQL.Port = os.Getenv("PGSQL_PORT")
		mc.PgSQL.Pass = os.Getenv("PGSQL_PASS")
		mc.PgSQL.Base = os.Getenv("PGSQL_BASE")
	}
	return *mc
}
