package model

import (
	"time"
)

type Todos struct {
	Id        int    `gorm:"column:id;primary key autoincrement" json:"id"`
	Title     string `gorm:"column:title" json:"title"`
	Done      int    `gorm:"column:done;type:int(1);default:0" json:"done"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// TableName 重命名表名
func (t *Todos) TableName() string {
	return "todos"
}

// AddOne 添加一条数据
func (t *Todos) AddOne() *Todos {
	db.Create(&t)
	return t
}

// FindOne 查询一条数据
func (t *Todos) FindOne(id int) *Todos {
	db.First(&t, "id=?", id)
	return t
}

// FindAll 查询所有数据
func (t *Todos) FindAll() (todos []Todos) {
	db.Find(&todos)
	return todos
}

// UpdateOne 修改一条数据
func (t *Todos) UpdateOne(id int) *Todos {
	db.Model(&t).Where("id=?", id).Updates(t)
	return t
}

// DeleteOne 删除一条数据
func (t *Todos) DeleteOne(id int) *Todos {
	db.Where("id=?", id).Delete(&t)
	return t
}
