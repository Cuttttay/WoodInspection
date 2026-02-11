package common

// 错误码定义
const (
	// 通用错误码
	CodeSuccess          = 0
	CodeInternalError    = 10001
	CodeInvalidParams    = 10002
	CodeUnauthorized     = 10003
	CodeForbidden        = 10004
	CodeNotFound         = 10005
	CodeMethodNotAllowed = 10006

	// 用户相关错误码
	CodeUserNotFound      = 20001
	CodeUserAlreadyExists = 20002
	CodeInvalidPassword   = 20003
	CodeInvalidToken      = 20004
	CodeCaptchaInvalid    = 20005

	// 检测模块错误码 (30000-30999)
	CodeDetectFileGetFailed        = 30001
	CodeDetectFileReadFailed       = 30002
	CodeDetectFailed               = 30003
	CodeDetectRecordQueryFailed    = 30004
	CodeDetectRecordNotFound       = 30005
	CodeDetectRecordIDInvalid      = 30006
	CodeDetectRecordSaveFailed     = 30007
	CodeDetectImageSaveFailed      = 30008
	CodeDetectImageInfoSaveFailed  = 30009
	CodeDetectImageOrURLRequired   = 30010
	CodeDetectVerdictInvalid       = 30011
	CodeDetectStartDateInvalid     = 30012
	CodeDetectEndDateInvalid       = 30013
	CodeDetectMinConfidenceInvalid = 30014
	CodeDetectStatisticsFailed     = 30015
	CodeDetectExportFailed         = 30016
	CodeDetectReportGetFailed      = 30017

	// 配置模块错误码 (40000-40999)
	CodeConfigNotInitialized  = 40001
	CodeConfigUpdateFailed    = 40002
	CodeThresholdUpdateFailed = 40003
	CodeModelUpdateFailed     = 40004

	// 图片上传错误码 (50000-50999)
	CodeImageUploadFailed       = 50001
	CodeImageSaveFailed         = 50002
	CodeImageInfoSaveFailed     = 50003
	CodeImageDefectDetectFailed = 50004

	// 缓存相关错误码
	CodeCacheError = 60001

	CodeJwcNotBound      = 70001
	CodeJwcLoginFailed   = 70002
	CodeJwcInvalidParams = 70003
	CodeJwcRequestFailed = 70004
	CodeJwcParseFailed   = 70005
)

// 错误信息映射
var errorMessages = map[int]string{
	CodeSuccess:           "成功",
	CodeInternalError:     "内部服务错误",
	CodeInvalidParams:     "参数错误",
	CodeUnauthorized:      "未授权",
	CodeForbidden:         "禁止访问",
	CodeNotFound:          "资源不存在",
	CodeMethodNotAllowed:  "只支持 POST 请求",
	CodeUserNotFound:      "用户不存在",
	CodeUserAlreadyExists: "用户已存在",
	CodeInvalidPassword:   "密码错误",
	CodeInvalidToken:      "令牌无效",
	CodeCacheError:        "缓存错误",
	CodeCaptchaInvalid:    "验证码错误",

	// 检测模块错误信息
	CodeDetectFileGetFailed:        "获取文件失败",
	CodeDetectFileReadFailed:       "读取文件失败",
	CodeDetectFailed:               "检测失败",
	CodeDetectRecordQueryFailed:    "查询检测记录失败",
	CodeDetectRecordNotFound:       "检测记录不存在",
	CodeDetectRecordIDInvalid:      "检测记录ID格式错误",
	CodeDetectRecordSaveFailed:     "保存检测记录失败",
	CodeDetectImageSaveFailed:      "保存图片信息失败",
	CodeDetectImageInfoSaveFailed:  "保存图片信息失败",
	CodeDetectImageOrURLRequired:   "必须提供图片文件或图片URL",
	CodeDetectVerdictInvalid:       "verdict 必须是 OK/NG/ALL",
	CodeDetectStartDateInvalid:     "start_date 格式必须是 YYYY-MM-DD",
	CodeDetectEndDateInvalid:       "end_date 格式必须是 YYYY-MM-DD",
	CodeDetectMinConfidenceInvalid: "min_confidence 必须在 0-1 之间",
	CodeDetectStatisticsFailed:     "获取统计信息失败",
	CodeDetectExportFailed:         "导出检测记录失败",
	CodeDetectReportGetFailed:      "获取报告失败",

	// 配置模块错误信息
	CodeConfigNotInitialized:  "配置未初始化",
	CodeConfigUpdateFailed:    "更新配置失败",
	CodeThresholdUpdateFailed: "更新阈值失败",
	CodeModelUpdateFailed:     "更新模型标识失败",

	// 图片上传错误信息
	CodeImageUploadFailed:       "上传图片失败",
	CodeImageSaveFailed:         "保存图片失败",
	CodeImageInfoSaveFailed:     "保存图片信息失败",
	CodeImageDefectDetectFailed: "缺陷检测失败",
}

// GetErrorMessage 获取错误信息
func GetErrorMessage(code int) string {
	if msg, ok := errorMessages[code]; ok {
		return msg
	}
	return "未知错误"
}

// AppError 应用错误
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// NewAppError 创建应用错误
func NewAppError(code int, message string) *AppError {
	if message == "" {
		message = GetErrorMessage(code)
	}
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// ErrorResponse 错误响应结构（用于直接返回JSON）
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(code int, customMessage ...string) ErrorResponse {
	message := GetErrorMessage(code)
	if len(customMessage) > 0 && customMessage[0] != "" {
		message = customMessage[0]
	}
	return ErrorResponse{
		Error: message,
	}
}

// NewErrorResponseWithDetail 创建带详细信息的错误响应
func NewErrorResponseWithDetail(code int, detail interface{}) map[string]interface{} {
	message := GetErrorMessage(code)
	return map[string]interface{}{
		"error":  message,
		"detail": detail,
	}
}
