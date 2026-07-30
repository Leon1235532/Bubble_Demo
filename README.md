<a id="readme-top"></a>

<div align="center">

<h1 align="center">GoTask</h1>

  <p align="center">
    基于 Gin + GORM + JWT 的待办任务管理系统，支持用户鉴权与回收站机制
    <br />
    <a href="https://github.com/Leon1235532/Bubble_Demo"><strong>Explore the docs »</strong></a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

GoTask 是一个采用 Handler-Service-Dao 三层架构的待办任务管理系统，提供完整的用户鉴权和 Todo CRUD 功能，支持软删除回收站机制。

核心功能：

1. 用户注册/登录 → JWT 签发 → 中间件鉴权
2. Todo 增删改查 + 软删除回收站
3. 回收站批量恢复与清空，按 uid 数据隔离

### Built With

![Tech Stack](https://skillicons.dev/icons?i=go,gin,mysql)

Gin · Gorm · JWT

<!-- GETTING STARTED -->

## Getting Started

### Prerequisites

- Go 1.21+
- MySQL 5.7+

### Installation

1. 克隆仓库
   ```sh
   git clone https://github.com/Leon1235532/Bubble_Demo.git
   ```
2. 配置数据库
   ```sh
   # 编辑 conf/config.ini，填入数据库连接信息
   ```
3. 启动服务
   ```sh
   go run main.go
   ```

服务将在 `http://localhost:8080` 启动。



<!-- USAGE -->
## Usage

### API 接口

| 模块 | 接口 | 说明 |
|------|------|------|
| 用户 | `POST /api/v1/register` | 注册（bcrypt 加密密码） |
| 用户 | `POST /api/v1/login` | 登录（返回 JWT Token） |
| 用户 | `PUT /api/v1/changePwd` | 修改密码 |
| 用户 | `DELETE /api/v1/user/:id` | 注销账号 |
| Todo | `GET /api/v1/todo?page=1&size=10` | 分页查询 Todo（需 Token） |
| Todo | `POST /api/v1/todo` | 新增 Todo |
| Todo | `PUT /api/v1/todo/:id` | 更新 Todo |
| Todo | `DELETE /api/v1/todo/:id` | 软删除（移入回收站） |
| 回收站 | `GET /api/v1/todo/recycle` | 查看回收站 |
| 回收站 | `PUT /api/v1/todo/recycle/:id` | 恢复单条 |
| 回收站 | `PUT /api/v1/todo/recycle` | 批量恢复 |
| 回收站 | `DELETE /api/v1/todo/recycle` | 清空回收站 |

### 项目结构

```
Bubble_Demo/
├── main.go                 # 入口
├── controller/             # Handler 层
├── service/                # Service 层
├── dao/                    # DAO 层（GORM 操作）
├── models/                 # 数据模型
├── middleware/             # Gin 中间件（JWT 鉴权）
├── conf/
│   └── config.ini          # 配置文件
├── .env                    # 敏感配置
└── pkg/
    └── e/                  # 响应码定义
```



<!-- CONTACT -->
## Contact

Leon1235532 - xzr12367@126.com

Project Link: [https://github.com/Leon1235532/Bubble_Demo](https://github.com/Leon1235532/Bubble_Demo)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
