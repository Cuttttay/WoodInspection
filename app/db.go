package app

import (
	"WoodInspection/internal/product/auth/model"
	model2 "WoodInspection/internal/product/dectect/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDBWithConfig 使用配置初始化数据库
func InitDBWithConfig(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User, config.Database.Pass, config.Database.Host, config.Database.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移 User 表结构（添加 role 字段等）
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, fmt.Errorf("数据库迁移失败: %w", err)
	}

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//启用日志以查看生成的SQL
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	//自动迁移所有模型（会自动创建索引）
	if err := db.AutoMigrate(
		&model.User{},
		&model2.Defect{},
	); err != nil {
		return nil, fmt.Errorf("数据库迁移失败")
	}
	return db, nil
}
