# lunar-gate

RBAC后台管理系统demo，已完成后端部分~~前端已加入todo list~~

## 项目结构

```bash
├── api/              # api
├── common
│   ├── query/        # 数据库查询封装
│   └── res/          # 响应封装
├── config/           # 全局配置
├── core/             # 初始化
├── docs/             # 文档 (TODO)
├── flags/            # 命令行参数
├── global/           # 全局变量
├── middleware/       # 中间件
├── model/            # 数据模型
├── routers/          # 路由
├── service           # 服务
│   └── redis_service/
├── utils/            # 工具类
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
└── settings.yaml     # 配置文件
```