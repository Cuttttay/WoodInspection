package service

import (
	"WoodInspection/internal/product/auth/common"
	"WoodInspection/internal/product/auth/model"
	"WoodInspection/internal/product/auth/repository"
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginResponse struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

type UserService interface {
	GetUserInfo(ctx context.Context, username string) (*model.User, error)
	UserLogin(ctx context.Context, username string, password string) (*UserLoginResponse, error)
}

type UserServiceImpl struct {
	userRepo  repository.UserRepository
	jwtSecret []byte
	jwtIssuer string
	jwtExpire time.Duration
}

func NewUserService(userRepo repository.UserRepository, jwtSecret []byte, jwtIssuer string) UserService {
	return &UserServiceImpl{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		jwtIssuer: jwtIssuer,
		jwtExpire: 168 * time.Hour,
	}
}

type Claims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func (c *UserServiceImpl) UserLogin(ctx context.Context, username string, password string) (*UserLoginResponse, error) {
	User, err := c.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, common.NewAppError(common.CodeUserNotFound, "用户不存在或密码错误")
	}
	if bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(password)) != nil {
		return nil, common.NewAppError(common.CodeUserNotFound, "用户不存在或密码错误")
	}

	claims := Claims{
		Id:   strconv.FormatInt(User.Id, 10),
		Name: User.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(c.jwtExpire)),
			Issuer:    c.jwtIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(c.jwtSecret)
	if err != nil {
		return nil, common.NewAppError(common.CodeInternalError, "生成令牌失败")
	}

	// 清除敏感信息
	userInfo := *User
	userInfo.Password = ""

	return &UserLoginResponse{
		Token: tokenString,
		User:  &userInfo,
	}, nil
}

func (c *UserServiceImpl) GetUserInfo(ctx context.Context, username string) (*model.User, error) {
	User, err := c.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, common.NewAppError(common.CodeUserNotFound, "")
	}
	User.Password = ""

	return User, nil

}
