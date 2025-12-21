package app

import (
	controller2 "WoodInspection/internal/product/auth/controller"
	"WoodInspection/internal/product/auth/repository"
	"fmt"

	service2 "WoodInspection/internal/product/auth/service"

	"gorm.io/gorm"
)

// Container 依赖注入容器
type Container struct {
	Config *Config
	DB     *gorm.DB

	// Repositories
	UserRepo repository.UserRepository

	UserService service2.UserService

	// Controllers
	UserController *controller2.UserController
}

// NewContainer 创建依赖注入容器
func NewContainer(configPath string) (*Container, error) {
	c := &Container{}

	// 加载配置
	if err := c.initConfig(configPath); err != nil {
		return nil, fmt.Errorf("初始化配置失败: %w", err)
	}

	// 初始化数据库
	if err := c.initDB(); err != nil {
		return nil, fmt.Errorf("初始化数据库失败: %w", err)
	}

	// 初始化 Repositories
	c.initRepositories()

	// 初始化 Services
	c.initServices()

	// 初始化 Controllers
	c.initControllers()

	return c, nil
}

// initConfig 初始化配置
func (c *Container) initConfig(configPath string) error {
	config, err := LoadConfigFromPath(configPath)
	if err != nil {
		return err
	}
	c.Config = config
	return nil
}

// initDB 初始化数据库
func (c *Container) initDB() error {
	db, err := InitDBWithConfig(c.Config)
	if err != nil {
		return err
	}
	c.DB = db
	return nil
}

// initRepositories 初始化 Repositories
func (c *Container) initRepositories() {
	c.UserRepo = repository.NewGormUserRepository(c.DB)
}

// initServices 初始化 Services
func (c *Container) initServices() {

	// User Service
	c.UserService = service2.NewUserService(
		c.UserRepo,
		[]byte(c.Config.JWT.Secret),
		c.Config.JWT.Issuer,
	)

}

// initControllers 初始化 Controllers
func (c *Container) initControllers() {
	c.UserController = controller2.NewUserController(c.UserService)
}

// Close 关闭资源
func (c *Container) Close() error {

	// 关闭数据库
	if c.DB != nil {
		sqlDB, err := c.DB.DB()
		if err != nil {
			return fmt.Errorf("获取数据库连接失败: %w", err)
		}
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("关闭数据库失败: %w", err)
		}
	}

	return nil
}
