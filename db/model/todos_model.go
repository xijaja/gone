package model

type Todos struct {
	Id    int    `gorm:"comment:id;type:int;primary key;autoincrement" json:"id"`
	Title string `gorm:"comment:标题;type:varchar(64);" json:"title"`
	Done  *bool  `gorm:"comment:完成;type:boolean;default:false" json:"done"`
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
	db.Order("id desc").Find(&todos)
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
