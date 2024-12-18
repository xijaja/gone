package model

// Todos 待办事项表
// 如果你的表模型中使用了外键，那么外键的逻辑字段不需要导出，使用小写且不需要 json Tag
// 注意：如果非要导出，那么查询返回值中将会多出不需要的字段，需要修改一些结构体方法
type Todos struct {
	Id    int    `gorm:"comment:id;type:bigint;primary key autoincrement;" json:"id"`
	Title string `gorm:"comment:标题;type:varchar(64);" json:"title"`
	Done  bool   `gorm:"comment:完成;type:boolean;default:false" json:"done"`
	// UsersID   int       `gorm:"comment:用户ID;type:bigint;" json:"users_id"` // 外键
	// user      Users     `gorm:"comment:用户;ForeignKey:UsersID;AssociationForeignKey:Id"` // 例如这个例子
	SomeTimesAt
}

// TableName 重命名表名
func (t *Todos) TableName() string {
	return "todos"
}

// AddOne 添加一条数据
func (t *Todos) AddOne() {
	// db.Create(&t)
	// 使用原生语句创建
	db.Exec("INSERT INTO todos (title, done) VALUES (?, ?)", t.Title, t.Done)
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
