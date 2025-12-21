# CORS 跨域配置指南

## CORS 预检请求（OPTIONS）

### 重要说明

前端在发送跨域请求时，浏览器会先发送一个 **OPTIONS** 预检请求，这是**正常行为**，不是错误。

### 为什么会有 OPTIONS 请求？

当满足以下条件时，浏览器会发送预检请求：
- 使用了自定义请求头（如 `Authorization`）
- `Content-Type` 不是简单类型（如 `application/json`）
- 使用了 `withCredentials: true`（携带 Cookie）

### 请求流程

1. **浏览器发送 OPTIONS 预检请求**
   ```
   OPTIONS /api/login
   Headers:
     Origin: http://localhost:5173
     Access-Control-Request-Method: POST
     Access-Control-Request-Headers: authorization,content-type
   ```

2. **后端必须正确响应 OPTIONS 请求**
   ```
   Status: 204 No Content
   Headers:
     Access-Control-Allow-Origin: http://localhost:5173
     Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE, PATCH
     Access-Control-Allow-Headers: Content-Type, Authorization, ...
     Access-Control-Allow-Credentials: true
     Access-Control-Max-Age: 86400
   ```

3. **浏览器收到允许后，发送真正的 POST 请求**
   ```
   POST /api/login
   Headers:
     Content-Type: application/json
     Authorization: Bearer {token}
     Origin: http://localhost:5173
   ```

### 后端处理 OPTIONS 请求

**当前实现**（Gin 框架）：
```go
// internal/middleware/cors.go
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 设置 CORS 响应头
        c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, ...")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, ...")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        
        // 处理 OPTIONS 预检请求
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204) // No Content
            return
        }
        
        c.Next()
    }
}
```

### 如何确认请求是否正常？

在浏览器开发者工具的 **Network** 标签中：

✅ **正常情况**：应该看到两个请求
1. `OPTIONS /api/login` (状态码 204)
2. `POST /api/login` (状态码 200)

❌ **异常情况**：只看到 OPTIONS 请求且失败
- 说明后端没有正确处理 OPTIONS 请求
- 需要检查后端的 CORS 配置
- 常见错误：
  - 没有设置 `Access-Control-Allow-Origin`
  - `Access-Control-Allow-Credentials: true` 时，`Access-Control-Allow-Origin` 不能是 `*`
  - 没有处理 OPTIONS 请求，直接返回 404

### CORS 响应头说明

| 响应头 | 说明 | 示例 |
|--------|------|------|
| `Access-Control-Allow-Origin` | 允许的源（域名） | `http://localhost:5173` |
| `Access-Control-Allow-Credentials` | 是否允许携带 Cookie | `true` |
| `Access-Control-Allow-Methods` | 允许的 HTTP 方法 | `POST, GET, OPTIONS` |
| `Access-Control-Allow-Headers` | 允许的请求头 | `Content-Type, Authorization` |
| `Access-Control-Max-Age` | 预检请求缓存时间（秒） | `86400` (24小时) |
| `Access-Control-Expose-Headers` | 暴露给前端的响应头 | `Content-Length, Content-Type` |

### 注意事项

1. **Credentials 和 Origin 的关系**
   - 当 `Access-Control-Allow-Credentials: true` 时
   - `Access-Control-Allow-Origin` **不能**设置为 `*`
   - 必须设置为具体的域名，如 `http://localhost:5173`

2. **预检请求缓存**
   - `Access-Control-Max-Age` 设置预检请求的缓存时间
   - 浏览器在缓存时间内不会重复发送预检请求
   - 建议设置为 86400 秒（24小时）

3. **开发环境 vs 生产环境**
   - 开发环境：可以允许所有源（但要注意 credentials 的限制）
   - 生产环境：必须明确指定允许的源，不能使用 `*`

### 当前配置的允许源

```go
allowedOrigins := []string{
    "http://localhost:3000",  // React 默认端口
    "http://localhost:5173",  // Vite 默认端口
    "http://localhost:8083",  // 后端端口
    "http://127.0.0.1:3000",
    "http://127.0.0.1:5173",
    "http://127.0.0.1:8083",
}
```

### 添加新的允许源

如果需要添加新的前端域名，修改 `internal/middleware/cors.go` 文件：

```go
allowedOrigins := []string{
    "http://localhost:3000",
    "http://localhost:5173",
    "http://localhost:8083",
    "http://127.0.0.1:3000",
    "http://127.0.0.1:5173",
    "http://127.0.0.1:8083",
    "https://your-production-domain.com", // 添加生产环境域名
}
```

### 常见问题排查

1. **OPTIONS 请求返回 404**
   - 检查路由是否正确配置
   - 确保 CORS 中间件在所有路由之前注册

2. **OPTIONS 请求返回 401**
   - OPTIONS 请求不应该经过认证中间件
   - 确保 CORS 中间件在认证中间件之前

3. **CORS 错误：Credentials 和 Origin 冲突**
   - 检查 `Access-Control-Allow-Origin` 是否为 `*`
   - 如果使用 credentials，必须设置为具体域名

4. **预检请求频繁发送**
   - 检查 `Access-Control-Max-Age` 是否设置
   - 确保浏览器正确缓存了预检请求结果


