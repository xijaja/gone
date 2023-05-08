package model

import (
	"fmt"
	"gone/db/access"
	"gone/start"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
)

// 数据库句柄
var db *gorm.DB

// 初始化数据库
func init() {
	// 初始化数据库
	db = access.NewConnect(start.Config.Database, access.Postgres).DB
	// 自动迁移，入参如 &Logs{}, &Todos{}
	err := db.AutoMigrate(sqlTagExecutor(&Logs{}, &Todos{})...)
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
}

// sqlTagExecutor sql tag 执行者
// 传入任意数量个结构体指针，获取结构体的 sql tag
// 如果不存在类型则创建，最后将指针返回用以自动迁移
func sqlTagExecutor(t ...interface{}) []interface{} {
	for _, v := range t {
		// 如果 v 是地址，则获取其指向的值
		if reflect.TypeOf(v).Kind() == reflect.Ptr {
			v = reflect.ValueOf(v).Elem().Interface()
		}
		// 读取所有字段的 sql tag
		for i := 0; i < reflect.TypeOf(v).NumField(); i++ {
			sql := reflect.TypeOf(v).Field(i).Tag.Get("sql")
			if len(sql) > 0 {
				// 将 sql tag 以空格拆分并获取第三个字段，例如：`create type phone_type as enum('android','iOS');"`
				tp := strings.Split(sql, " ")[2]
				haveTp := db.Exec(fmt.Sprintf("select 1 from pg_type where typname = '%s';", tp))
				// 如果不存在该枚举类型，则创建
				if haveTp.RowsAffected == 0 {
					// 如果 sql 最后一个字符不是分号，则加上分号
					if sql[len(sql)-1] != ';' {
						sql += ";"
					}
					db.Exec(sql)
					fmt.Println(fmt.Sprintf("初始化创建 %s 类型，执行 sql 语句: %s", tp, sql))
				}
			}
		}
	}
	return t
}
