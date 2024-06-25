package access

import (
	"embed"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

//go:embed db.sqlite
var LiteDB embed.FS

// Connect 数据库连接
type Connect struct {
	DB *gorm.DB
}

// NewConnect 初始化数据库连接
func NewConnect() *Connect {
	return &Connect{
		DB: initSqlite(),
	}
}

// 初始化 Sqlite 数据库
func initSqlite() *gorm.DB {
	// 使用嵌入的 sqlite 数据库
	dbFile, err := LiteDB.ReadFile("db.sqlite")
	if err != nil {
		log.Fatal("读取 db.sqlite 文件失败")
	}
	// 数据库连接
	var db *gorm.DB
	// 使用相对路径的 sqlite 数据库
	// var db, dbErr = gorm.Open(sqlite.Open("db/access/db.sqlite"), &gorm.Config{})
	// 使用内存的 sqlite 数据库
	db, dbErr := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	// 数据库连接错误
	if dbErr != nil {
		panic("链接 Sqlite 数据库恐慌：" + dbErr.Error())
	}
	// 执行数据库文件，初始化数据库
	if err := db.Exec(string(dbFile)).Error; err != nil {
		panic("初始化 Sqlite 数据库恐慌：" + err.Error())
	}
	return db
}
