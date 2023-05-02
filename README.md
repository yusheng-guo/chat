# chat
Chat system developed using Gin and github.com/gobwas/ws!

## Project Design
#### Build Project Catalog Structure
```shell
chat
├─api    # 接口
│  └─v1  # 版本
├─assets # 静态资源
│  ├─images      # 图片
│  └─screenshots # 截图
├─client   # 简单的客户端实现
├─config   # 配置处理
├─docs     # swagger生成的接口文档
├─internal # 内部文件
│  ├─dao     # 数据访问层
│  ├─middleware # 中间件
│  ├─model   # 模型
│  ├─routers # 路由
│  └─service # 服务
├─log    # 日志处理
├─tmp    # 临时文件
├─utils  # 工具函数
└─vendor # 项目以来的第三方库
```
#### Create Database
- Database: chat
- Charts: users, messages


#### Create Model
- Model
- User
- Message


#### Routing
- 用户管理

| Function | Mathod | Path         | Remark             |
|----------|--------|--------------|--------------------|
| 用户注册     | POST   | /register    |                    |
| 用户登录     | POST   | /login       |                    |
| 管理员登录    | POST   | /admin       | IsActive(默认为false) |
| 用户注销     | POST   | /logout      |                    |
| 修改用户信息   | PUT    | /me          |                    |
| 获取用户信息   | GET    | /me          |                    |

- 朋友管理

| Function | Mathod | Path         | Remark             |
|----------|--------|--------------|--------------------|
| 添加好友     | POST   | /friend/:id  |                    |
| 删除好友     | DELETE | /friend/:id  |                    |
| 更改好友信息   | PUT    | /friend/:id  |                    |
| 获取好友信息   | GET    | /friend/:id  |                    |
| 获取所有好友   | GET    | /friends     |                    |

- 消息管理

| Function | Mathod | Path         | Remark             |
|----------|--------|--------------|--------------------|
| 发送消息     | POST   | /message     |                    |
| 删除消息     | DELETE | /message/:id |                    |
| 修改消息     | PUT    | /message/:id |                    |
| 获取消息     | GET    | /message     |                    |


## Public Components
#### Configuration Management [`viper`](https://github.com/spf13/viper)
    
#### Database Connection 
- [`rethinkdb`](https://rethinkdb.com/)
- [`Redis`](https://redis.io/)

#### Response Processing 
- Response
- ErrorResponse

## API Documentation [`swagger`](https://swagger.io/)
![image](https://github.com/yushengguo557/chat/blob/main/assets/screenshots/swagger.png)

## Function Development
#### Register and Login 

#### WebSocket Communication 