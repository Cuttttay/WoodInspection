package dto

import "os"

type DetectOne struct {
	file os.FileInfo
}

// SaveDetectRecordRequest 保存检测记录请求
type SaveDetectRecordRequest struct {
	ImageURL string                   `json:"image_url,omitempty" binding:"omitempty"` // 图片URL（可选）
	Defects  []map[string]interface{} `json:"defects" binding:"required"`              // 缺陷列表（必填）
	Verdict  string                   `json:"verdict,omitempty"`                       // 判定结果：OK/NG
	Model    string                   `json:"model,omitempty"`                         // 使用的模型名称
}

// GetRecordsQuery 检测记录查询参数
// 对应文档 3.1：
// - page
// - page_size
// - start_date
// - end_date
// - verdict
// - min_confidence
// - model_name
type GetRecordsQuery struct {
	Page          int     `form:"page"`
	PageSize      int     `form:"page_size"`
	StartDate     string  `form:"start_date"`
	EndDate       string  `form:"end_date"`
	Verdict       string  `form:"verdict"`
	MinConfidence float64 `form:"min_confidence"`
	ModelName     string  `form:"model_name"`
}
