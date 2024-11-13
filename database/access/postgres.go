package access

import (
	"fmt"
	"gone/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// InitPostgresSQL 初始化 PostgresSQL 数据库
func InitPostgresSQL(host, user, port, pass, base, sslmode string) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		host, user, pass, base, port, sslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 newLogger, // Gorm SQL 日志全局模式
		SkipDefaultTransaction: true,      // 禁用默认事务，提升性能
		PrepareStmt:            true,      // 执行 SQL 时缓存，提高调用速度
	})
	if err != nil {
		panic(fmt.Sprintf("致命错误：无法连接到数据库: %v", err))
	}
	return db
}

// // InitMysql 初始化 MySQL 数据库 (如果你使用 mysql 数据库，可以使用这个方法)
// func InitMysql(my start.Database) (db *gorm.DB) {
// 	// 拼接数据库连接信息
// 	dsn := fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		my.User, my.Pass, my.Host, my.Port, my.Base,
// 	)
// 	// 初始化db
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		// DisableForeignKeyConstraintWhenMigrating: true,      // 禁用创建外键约束
// 		Logger:                 newLogger, // Gorm SQL 日志全局模式
// 		SkipDefaultTransaction: true,      // 禁用默认事务，提升性能
// 		PrepareStmt:            true,      // 执行 SQL 时缓存，提高调用速度
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
		LogLevel:                  getLogLevel(), // 日志级别
		IgnoreRecordNotFoundError: true,          // 忽略 ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,         // 禁用彩色打印
	},
)

// 确定日志级别
func getLogLevel() logger.LogLevel {
	if *config.S {
		return logger.Silent
	}
	return logger.Info
}
