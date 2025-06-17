# lunar-gate

RBAC后台管理系统demo

[![Go Report Card](https://goreportcard.com/badge/github.com/Xavier-Tse/lunar-gate)](https://goreportcard.com/report/github.com/Xavier-Tse/lunar-gate)

## 项目结构

```bash
├── api/              # api
├── common
│   ├── query/        # 数据库查询封装
│   └── res/          # 响应封装
├── config/           # 全局配置
├── core/             # 初始化
├── flags/            # 命令行参数
├── global/           # 全局变量
├── middleware/       # 中间件
├── model/            # 数据模型
├── routers/          # 路由
├── service           # 服务
│   └── redis_service/
├── utils/            # 工具类
├── web/              # 前端
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
└── settings.yaml     # 配置文件
```

## 使用方法

1. 填写配置文件`settings.yaml`
2. 运行
    ```bash
   # 前端
   cd web
   npm install
   npm run dev
   
   # 后端
   go mod download
   go run main.go
   ```

## 支持的命令行参数

```bash
go run main.go -db # 自动迁移数据库
go run main.go -f "foo.yaml" # 指定配置文件
go run main.go -m user -t create # 创建管理员用户
```