# Wood Inspection System

一个基于AI的木材缺陷检测系统，能够自动识别和分析木材表面的缺陷，提供准确的质量评估和报告。

## 项目概述

Wood Inspection System 是一个企业级的木材质量检测解决方案，结合了先进的计算机视觉技术和深度学习算法。系统能够：

- 🔍 **自动缺陷检测**: 使用YOLOv12模型自动检测木材表面缺陷
- 📊 **智能质量评估**: 基于置信度阈值和缺陷数量进行OK/NG判定
- 📈 **数据统计分析**: 提供详细的检测统计和趋势分析
- 🔐 **用户认证系统**: 基于JWT的安全认证机制
- 📁 **记录管理**: 完整的检测记录存储和查询功能
- 📤 **数据导出**: 支持检测记录的批量导出

## 技术栈

### 后端框架
- **Go 1.25**: 主要编程语言
- **Gin**: HTTP Web框架，提供RESTful API
- **GORM**: ORM框架，用于数据库操作
- **MySQL**: 关系型数据库

### 核心依赖
- **JWT**: 用户认证和授权
- **Viper**: 配置管理
- **bcrypt**: 密码加密

### AI模型
- **YOLOv12**: 木材缺陷检测模型
- **ngrok**: 模型API代理服务

## 系统架构

```
WoodInspection/
├── api/                    # API路由配置
├── app/                    # 应用容器和配置
├── config/                 # 配置文件
├── internal/
│   ├── middleware/         # 中间件（CORS, JWT认证）
│   └── product/
│       ├── auth/          # 用户认证模块
│       ├── dectect/       # 缺陷检测模块
│       ├── image/         # 图片处理模块
│       ├── WoodModel/     # AI模型集成
│       └── config/        # 配置管理模块
└── main.go                # 应用入口
```

## 核心功能

### 1. 缺陷检测API
- **POST /api/detect**: 上传图片进行缺陷检测
- 支持多种图片格式
- 返回缺陷位置、类型和置信度

### 2. 检测记录管理
- **POST /api/detect/records**: 保存检测记录
- **GET /api/detect/records**: 获取检测记录列表（支持分页、筛选）
- **GET /api/detect/records/:id**: 获取单条记录详情
- **GET /api/detect/records/export**: 导出检测记录

### 3. 统计分析
- **GET /api/detect/statistics**: 获取检测统计信息
  - 总检测数量
  - OK/NG比率
  - 缺陷类型分布
  - 置信度分布

### 4. 配置管理
- **GET /api/config**: 获取系统配置
- **PUT /api/config**: 更新系统配置
- **GET /api/config/threshold**: 获取置信度阈值
- **PUT /api/config/threshold**: 更新置信度阈值
- **GET /api/models**: 获取模型列表
- **PUT /api/models/current**: 切换当前模型

### 5. 用户认证
- **POST /api/login**: 用户登录
- **GET /api/user/info**: 获取用户信息（需要认证）
