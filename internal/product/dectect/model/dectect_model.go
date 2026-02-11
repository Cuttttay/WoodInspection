package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Bbox 边界框类型，用于存储 [x1, y1, x2, y2]
type Bbox [4]int

// Value 实现 driver.Valuer 接口，用于存储到数据库
func (b Bbox) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// Scan 实现 sql.Scanner 接口，用于从数据库读取
func (b *Bbox) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, b)
}

type Defect struct {
	ID int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"` // 可选
	//添加索引ImageID最常查询
	ImageID int64   `gorm:"column:image_id;index:idx_image_id" json:"image_id,omitempty"` // 如果确实需要关联
	Label   string  `gorm:"type:varchar(100);not null;column:label" json:"label"`         // 必填
	Score   float64 `gorm:"type:decimal(5,4);not null;column:score" json:"score"`         // 必填

	X    *int `gorm:"column:x" json:"x,omitempty"`
	Y    *int `gorm:"column:y" json:"y,omitempty"`
	Area *int `gorm:"column:area" json:"area,omitempty"`
	//复合索引,(image_id,level)用于查询某张图片的特定等级缺陷
	Level *string `gorm:"type:varchar(50);column:level;index:idx_image_level,priority:2" json:"level,omitempty"`

	Bbox Bbox `gorm:"type:json;column:bbox" json:"bbox"` // 必填：[x1,y1,x2,y2]
}

// TableName 指定表名
func (Defect) TableName() string {
	return "defects"
}

// DefectReport 缺陷检测报告
type DefectReport struct {
	ID int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	//外键+索引
	ImageID int64 `gorm:"column:image_id;not null;index:idx_image_id" json:"image_id"`
	//判定结果加索引
	Verdict string `gorm:"type:enum('OK','NG','UNKNOWN');default:'UNKNOWN';column:verdict;index:idx_vwedict_time,priority:1" json:"verdict"`
	//缺陷数量索引
	DefectsCount int   `gorm:"column:defects_count;index:idx_verdict_count;default:0" json:"defects_count"`
	ReportData   JSONB `gorm:"type:json;column:report_data" json:"report_data"`
	//时间索引
	CreatedAt time.Time `gorm:"type:datetime;index:idx_verdict_time,priority:2;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`
}

// TableName 指定表名
func (DefectReport) TableName() string {
	return "defect_reports"
}

// JSONB 自定义 JSON 类型，用于存储 JSON 数据
type JSONB map[string]interface{}

// Value 实现 driver.Valuer 接口
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan 实现 sql.Scanner 接口
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}
