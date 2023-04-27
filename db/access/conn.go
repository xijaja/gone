package access

import (
	"embed"
	"fmt"
	"gone/start"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:embed sqlite.db
var LiteDB embed.FS

// DB 数据库连接
var DB *gorm.DB

func init() {
	// 读取配置文件
	config := start.Config
	// 初始化数据库
	if config.UsePgSQL {
		DB = initPostgresSQL()
	} else {
		DB = initSqlite()
	}
	// 自动迁移，入参如 &Movie{}, &Todos{}
	if err := DB.AutoMigrate(); err != nil {
		fmt.Println("数据库迁移失败:", err)
	}
	fmt.Println("数据库初始化完成")
}

// 初始化 Sqlite 数据库
func initSqlite() *gorm.DB {
	lite, _ := LiteDB.Open("lite.db")                                   // 获取 LiteDB 的 lite.db 文件
	liteDB, _ := lite.Stat()                                            // 获取 lite.db 文件的信息
	var db, err = gorm.Open(sqlite.Open(liteDB.Name()), &gorm.Config{}) // 使用嵌入的 sqlite 数据库
	// var db, err = gorm.Open(sqlite.Open("lite.db"), &gorm.Config{}) // 使用相对路径的 sqlite 数据库
	if err != nil {
		panic("初始化 Sqlite 数据库恐慌：" + err.Error())
	}
	return db
}

// 初始化 PostgresSQL 数据库
func initPostgresSQL() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,      // 禁用创建外键约束
		Logger:                                   newLogger, // Gorm SQL 日志全局模式
		SkipDefaultTransaction:                   true,      // 禁用默认事务，提升性能
		PrepareStmt:                              true,      // 执行 SQL 时缓存，提高调用速度
	})
	if err != nil {
		panic("初始化 PostgreSQL 数据库恐慌：" + err.Error())
	}
	return db
}

// // 初始化 MySQL 数据库 (如果你使用 mysql 数据库，可以使用这个方法)
// func initMysql(user, pwd, addr, port, base string) (db *gorm.DB) {
// 	// 拼接数据库连接信息
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, addr, port, base)
// 	// 初始化db
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		DisableForeignKeyConstraintWhenMigrating: true,      // 禁用创建外键约束
// 		Logger:                                   newLogger, // Gorm SQL 日志全局模式
// 		SkipDefaultTransaction:                   true,      // 禁用默认事务，提升性能
// 		PrepareStmt:                              true,      // 执行 SQL 时缓存，提高调用速度
// 	})
// 	if err != nil {
// 		panic("初始化 MySQL 数据库恐慌：" + err.Error())
// 	}
// 	return db
// }

// Gorm SQL 日志全局模式
var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer 日志输出的目标，前缀和日志包含的内容
	logger.Config{
		SlowThreshold:             time.Second,   // 慢 SQL 阈值
		LogLevel:                  logger.Silent, // 日志级别
		IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,         // 禁用彩色打印
	},
)
