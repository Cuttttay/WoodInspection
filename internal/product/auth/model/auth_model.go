package model

import "time"

type User struct {
	Id int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	//添加唯一索引：username
	Username string `gorm:"type:varchar(100);not null;column:username" json:"name"`
	Password string `gorm:"type:varchar(255);not null;column:password" json:"password"`
	//添加普通索引，按角色查询用户
	Role string `gorm:"type:varchar(50);default:'user';idx_role;column:role" json:"role"`
	//添加索引：需要按创建时间顺序排列时
	CreatedTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;index:idx_created_time;column:created_time" json:"created_time"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
