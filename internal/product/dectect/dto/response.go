package dto

import "time"

// PredictResponse API 预测响应
type PredictResponse struct {
	Model      string        `json:"model,omitempty"`      // 模型名称，如 "yolov12"
	Count      int           `json:"count,omitempty"`      // 检测到的缺陷数量
	Detections []interface{} `json:"detections,omitempty"` // 检测结果数组（实际 API 返回的字段）

	// 兼容字段（如果 API 返回不同格式）
	Image       string        `json:"image,omitempty"`
	Defect      interface{}   `json:"defect,omitempty"`
	Verdict     string        `json:"verdict,omitempty"`
	Predictions []interface{} `json:"predictions,omitempty"` // 兼容旧格式
	Confidence  float64       `json:"confidence,omitempty"`
}

type DetectResponse struct {
	Image   string    `json:"image"`
	Defect  DetectOne `json:"defect"`
	Verdict string    `json:"verdict"`
}

type StatisticsResponse struct {
	TotalRecords           int64              `json:"total_records"`
	OkCount                int64              `json:"ok_count"`
	NgCount                int64              `json:"ng_count"`
	NgRate                 float64            `json:"ng_rate"`
	DefectTypes            map[string]int64   `json:"defect_types"`            // 缺陷类型计数
	ConfidenceDistribution map[string]int64   `json:"confidence_distribution"` // 置信度分布分桶
	DateRange              *DateRangeResponse `json:"date_range,omitempty"`    // 起止日期（可选）
}

type DateRangeResponse struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type GetStatisticsQuery struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	ModelName string `form:"model_name"`
}

// DetectRecordItem 单条检测记录（列表中的一项）
type DetectRecordItem struct {
	ID                  int       `json:"id"`
	ImageID             int       `json:"image_id"`
	ImageURL            string    `json:"image_url"`
	Verdict             string    `json:"verdict"`
	DefectsCount        int       `json:"defects_count"`
	Model               string    `json:"model"`
	ConfidenceThreshold float64   `json:"confidence_threshold"`
	CreatedAt           time.Time `json:"created_at"`
}

// DetectRecordListResponse 检测记录列表响应（对应文档 3.1）
type DetectRecordListResponse struct {
	Total      int                `json:"total"`
	Page       int                `json:"page"`
	PageSize   int                `json:"page_size"`
	TotalPages int                `json:"total_pages"`
	Records    []DetectRecordItem `json:"records"`
}

// DetectRecordDetailResponse 检测记录详情响应（对应文档 3.2）
type DetectRecordDetailResponse struct {
	ID                  int         `json:"id"`
	ImageID             int         `json:"image_id"`
	ImageURL            string      `json:"image_url"`
	FilePath            string      `json:"file_path"`
	Verdict             string      `json:"verdict"`
	DefectsCount        int         `json:"defects_count"`
	Model               string      `json:"model"`
	ConfidenceThreshold float64     `json:"confidence_threshold"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	ReportData          interface{} `json:"report_data"` // 里面包含 defects 等详细信息
}

type DetectRecordQuery struct {
	Page, PageSize     int
	StartDate, EndDate string
	Verdict            string // OK/NG/ALL
	MinConfidence      float64
	ModelName          string
}

type DetectRecordListOne struct {
	Id                  int       `json:"id"`
	ImageId             int       `json:"image_id"`
	ImageURL            string    `json:"image_url"`
	FilePath            string    `json:"file_path"`
	Verdict             string    `json:"verdict"`
	DefectsCount        int       `json:"defects_count"`
	Model               string    `json:"model"`
	ConfidenceThreshold float64   `json:"confidence_threshold"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Defects             []struct{}
}

type DefectTypes struct {
	DeadKnot    int
	LiveKnot    int
	Crack       int
	ResinPocket int
}
