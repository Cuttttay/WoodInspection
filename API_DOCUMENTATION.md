# 前端对接 API 文档

## 基础信息

### 服务器地址
- **开发环境**: `http://localhost:8083`
- **生产环境**: 根据实际部署情况修改

### 端口
- **后端端口**: `8083` (可在 `config/config.yaml` 中修改)

### 基础 URL
```
http://localhost:8083/api
```

---

## 统一响应格式

### 成功响应
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    // 具体数据
  }
}
```

### 错误响应
```json
{
  "code": 10002,
  "message": "参数错误"
}
```

### 错误码说明
| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 10001 | 内部服务错误 |
| 10002 | 参数错误 |
| 10003 | 未授权 |
| 10004 | 禁止访问 |
| 10005 | 资源不存在 |
| 20001 | 用户不存在 |
| 20002 | 用户已存在 |
| 20003 | 密码错误 |
| 20004 | 令牌无效 |

---

## API 接口列表

### 1. 用户登录

**接口地址**: `POST /api/login`

**请求头**:
```
Content-Type: application/json
```

**请求体**:
```json
{
  "username": "string",
  "password": "string"
}
```

**响应示例** (成功):
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "用户名",
      "role": "user",
      "created_time": "2024-01-01T00:00:00Z"
    }
  }
}
```

**响应示例** (失败):
```json
{
  "code": 20001,
  "message": "用户不存在或密码错误"
}
```

**Cookie 说明**:
- 登录成功后，服务器会自动设置 Cookie: `access_token`
- Cookie 有效期: 168 小时 (7天)
- Cookie 属性: HttpOnly, Secure

**前端处理方式**:
- 方式一: 使用 Cookie（推荐，自动携带，无需手动处理）
- 方式二: 从响应中获取 `token`，后续请求在 Header 中携带: `Authorization: Bearer {token}`

---

### 2. 获取当前用户信息

**接口地址**: `GET /api/user/info`

**认证要求**: 需要 JWT 认证

**请求头**:
```
Authorization: Bearer {token}
```
或者使用 Cookie（如果登录时设置了 Cookie，会自动携带，无需手动添加）

**响应示例** (成功):
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "id": 1,
    "name": "用户名",
    "role": "user",
    "created_time": "2024-01-01T00:00:00Z"
  }
}
```

**响应示例** (未授权):
```json
{
  "code": 10003,
  "message": "未授权"
}
```

---

## 前端配置示例

### Axios 配置 (Vue/React)

```javascript
import axios from 'axios'

// 创建 axios 实例
const api = axios.create({
  baseURL: 'http://localhost:8083/api',
  timeout: 10000,
  withCredentials: true, // 允许携带 Cookie
})

// 请求拦截器 - 添加 token
api.interceptors.request.use(
  config => {
    // 方式一: 从 Cookie 自动携带（推荐）
    // 如果使用 Cookie，不需要手动添加 token
    
    // 方式二: 从 localStorage 获取 token
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器 - 处理错误
api.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code === 0) {
      return res.data
    } else {
      // 处理业务错误
      console.error('业务错误:', res.message)
      return Promise.reject(new Error(res.message))
    }
  },
  error => {
    // 处理 HTTP 错误
    if (error.response?.status === 401) {
      // 未授权，跳转到登录页
      // router.push('/login')
    }
    return Promise.reject(error)
  }
)

export default api
```

### 使用示例

```javascript
// 登录
async function login(username, password) {
  try {
    const response = await api.post('/login', {
      username,
      password
    })
    
    // 如果使用 token 方式，保存到 localStorage
    if (response.token) {
      localStorage.setItem('token', response.token)
    }
    
    // 如果使用 Cookie 方式，token 会自动保存
    return response
  } catch (error) {
    console.error('登录失败:', error.message)
    throw error
  }
}

// 获取用户信息
async function getUserInfo() {
  try {
    const userInfo = await api.get('/user/info')
    return userInfo
  } catch (error) {
    console.error('获取用户信息失败:', error.message)
    throw error
  }
}
```

### Fetch 配置

```javascript
const API_BASE_URL = 'http://localhost:8083/api'

// 通用请求函数
async function request(url, options = {}) {
  const token = localStorage.getItem('token')
  
  const config = {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...(token && { Authorization: `Bearer ${token}` }),
      ...options.headers,
    },
    credentials: 'include', // 允许携带 Cookie
  }
  
  const response = await fetch(`${API_BASE_URL}${url}`, config)
  const data = await response.json()
  
  if (data.code === 0) {
    return data.data
  } else {
    throw new Error(data.message)
  }
}

// 登录
async function login(username, password) {
  const data = await request('/login', {
    method: 'POST',
    body: JSON.stringify({ username, password }),
  })
  
  if (data.token) {
    localStorage.setItem('token', data.token)
  }
  
  return data
}

// 获取用户信息
async function getUserInfo() {
  return await request('/user/info')
}
```

---

## CORS 配置

后端已配置 CORS，允许以下源访问：
- `http://localhost:3000` (React 默认端口)
- `http://localhost:5173` (Vite 默认端口)
- `http://localhost:8083`
- `http://127.0.0.1:3000`
- `http://127.0.0.1:5173`
- `http://127.0.0.1:8083`

如果需要添加其他源，请修改 `internal/middleware/cors.go` 文件。

### OPTIONS 预检请求

**重要**：前端在发送跨域请求时，浏览器会先发送一个 **OPTIONS** 预检请求，这是**正常行为**，不是错误。

**请求流程**：
1. 浏览器发送 `OPTIONS /api/login` 预检请求
2. 后端返回 204 状态码，并设置 CORS 响应头
3. 浏览器收到允许后，发送真正的 `POST /api/login` 请求

**在浏览器开发者工具的 Network 标签中**：
- ✅ **正常情况**：应该看到两个请求
  1. `OPTIONS /api/login` (状态码 204)
  2. `POST /api/login` (状态码 200)
- ❌ **异常情况**：只看到 OPTIONS 请求且失败，说明后端 CORS 配置有问题

**详细说明**：请查看 `docs/CORS_GUIDE.md` 文档。

---

## 注意事项

1. **Token 存储**: 
   - 推荐使用 Cookie（HttpOnly，更安全）
   - 或使用 localStorage（需要手动在请求头添加）

2. **请求头**: 
   - 所有 POST 请求需要设置 `Content-Type: application/json`
   - 需要认证的接口需要携带 `Authorization: Bearer {token}` 或使用 Cookie

3. **错误处理**: 
   - 统一检查响应中的 `code` 字段
   - `code === 0` 表示成功
   - 其他 `code` 值表示错误，查看 `message` 获取错误信息

4. **跨域和 OPTIONS 请求**: 
   - 开发环境已配置 CORS，支持跨域请求
   - 浏览器会自动发送 OPTIONS 预检请求，这是正常行为
   - 如果看到 OPTIONS 请求失败，检查后端 CORS 配置
   - 生产环境需要根据实际情况配置允许的源

5. **Cookie 和 CORS**: 
   - 使用 Cookie 时，前端必须设置 `withCredentials: true`（Axios）或 `credentials: 'include'`（Fetch）
   - 后端已配置 `Access-Control-Allow-Credentials: true`
   - 预检请求会被缓存 24 小时（`Access-Control-Max-Age: 86400`），减少重复请求

