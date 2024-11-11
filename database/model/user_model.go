package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRole 用户权限
type UserRole string

// 用户权限枚举
const (
	NormalUserRole UserRole = "normal" // 普通用户
	AdminUserRole  UserRole = "admin"  // 管理员
	SuperUserRole  UserRole = "super"  // 超级管理员
)

// User 用户表
type User struct {
	Id       uuid.UUID `gorm:"comment:id;type:uuid;primary key;" json:"id"`
	Username string    `gorm:"comment:登录名;unique;" json:"username"`
	Password string    `gorm:"comment:密码MD5值;" json:"password"`
	Role     UserRole  `sql:"create type user_role as enum('normal','admin','super');"  gorm:"comment:权限;default:normal" json:"role"`
	SomeTimesAt
}

// BeforeCreate 在创建前调用，用于生成 uuid
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Id = uuid.New()
	return
}

func (u *User) FindOneByUsername(username string) *gorm.DB {
	return db.First(&u, "username = ? AND deleted_at is null", username)
}

func (u *User) CreateOne() *gorm.DB {
	return db.Create(&u)
}
