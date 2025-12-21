// Package handler handler/auth_handler.go
package handler

import (
	"WoodInspection/internal/product/auth/common"
	"WoodInspection/internal/product/auth/dto"
	"WoodInspection/internal/product/auth/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userSvc service.UserService
}

func NewAuthHandler(userSvc service.UserService) *AuthHandler {
	return &AuthHandler{userSvc: userSvc}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    common.CodeInvalidParams, // 你按项目已有 code 来
			"message": "请求参数格式错误",
			"data":    nil,
		})
		return
	}

	resp, err := h.userSvc.UserLogin(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		// 如果 err 是你自定义 AppError，可以取 code/message
		var appErr *common.AppError
		if errors.As(err, &appErr) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"data":    nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    common.CodeInternalError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    resp, // ✅ 这里才会让前端 res.data 不是空
	})
}
