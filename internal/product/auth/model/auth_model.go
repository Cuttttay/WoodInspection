package model

import "time"

type User struct {
	Id          int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Username    string    `gorm:"type:varchar(100);not null;column:username" json:"name"`
	Password    string    `gorm:"type:varchar(255);not null;column:password" json:"password"`
	Role        string    `gorm:"type:varchar(50);default:'user';column:role" json:"role"`
	CreatedTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_time" json:"created_time"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
