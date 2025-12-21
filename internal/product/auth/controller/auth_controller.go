package controller

import (
	common "WoodInspection/internal/product/auth/common"
	"WoodInspection/internal/product/auth/dto"
	"WoodInspection/internal/product/auth/service"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc service.UserService
}

func NewUserController(userSvc service.UserService) *UserController {
	return &UserController{userSvc: userSvc}
}

func (u *UserController) UserLogin(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		common.ErrorWithAppError(c, common.NewAppError(common.CodeInvalidParams, ""))
		return
	}

	loginResp, err := u.userSvc.UserLogin(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if appErr, ok := err.(*common.AppError); ok {
			common.ErrorWithAppError(c, appErr)
		} else {
			fmt.Println(err)
		}
		return
	}

	maxAge := int((168 * time.Hour).Seconds())
	c.SetCookie("access_token", loginResp.Token, maxAge, "/", "", true, true)

	common.Success(c, loginResp)

}

func (u *UserController) GetUserInfo(c *gin.Context) {
	uid, ok := c.Get("id")
	if !ok {
		common.Error(c, common.CodeUnauthorized, "unauthorized")
		return
	}

	user, err := u.userSvc.GetUserInfo(c.Request.Context(), uid.(string))
	if err != nil {
		if appErr, ok := err.(*common.AppError); ok {
			common.ErrorWithAppError(c, appErr)
		} else {
			common.Error(c, common.CodeInternalError, "internal error")
		}
		return
	}

	common.Success(c, user)
	
}
