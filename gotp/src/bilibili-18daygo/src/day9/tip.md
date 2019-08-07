1. 建立连接，传输数据
## socket编程

### 服务端处理流程
1. 监听端口
2. 接收客户端连接
3. 创建goroutine，处理连接
### 客户端处理流程
1. 建立与服务端的连接
2. 进行数据收发
3. 关闭连接

## redis
[redigo](https://github.com/gomodule/redigo)
开源的高性能key-value内存数据库
- 集合操作: Set
- 哈希表: HSet
- 批量Set: Mset
- 队列操作: lpush

### redis连接池
- 过期时间: expire

### 
1. 海量用户在线聊天
2. 点对点聊天
3. 用户登录&注册

#### 服务端开发
1. 用户管理
- 用户id：数字
- 用户密码：字母数字组合
- 用户昵称：用来显示
- 用户性别：字符串
- 用户头像：url
- 用户上线登录时间：字符串
- 用户是否在线：online
- 数据存储：redis hash：users
2. 用户动作
- 发送消息
- 接收消息
-
3. 用户注册&登陆

4. 用户消息离线存储

#### 客户端开发
1. 用户注册
2. 用户登录
3. 发送消息
4. 获取用户列表

#### 通信协议
[0:4]表示长度
[]json

104 

