 # go-ums 模型设计( v0.1.0 )

go-ums 的用户模型设计



## 0. 说明
这是一个从零开始的练手项目, 持续渐进, 步步简洁吧

## 1. 业务场景

从简单的业务场景开始, 

### 1.1 v0.1.0

用户相关操作

 - [ ] 用户注册 register 
    用户以 email + 密码 注册提交个人信息, 提交后验证 email 是否存在唯一冲突, 如果没有, 注册成功( 分配用户ID)
注意, 这里隐藏着一个检查用户是否存在的操作
 - [ ] 用户登录与登出 login / logout 
    用户以  email + passwrd 或 用户名 + password 或 userID + password 可以进行登录, 登录成功, 给用户发送通行令牌 accessToken 
    注意, 这里隐藏着一个检查用户登录成功返回一个 access token 的操作
 - [ ] 用户认证 authentication
    针对具体业务, 用户访问具体业务时, 发送 access token 进行认证, 如果用户认证成功, 可以访问指定业务, 否则拒绝. 一般来说, 用户认证成功即表示创建一个合法的业务会话
 - [ ] 用户通行令牌验证 verify
    用户访问指定业务时, 验证 token 是否在有效期内, 以及相关授权是否有变更, 如不符合业务授权限制, 则拒绝提供服务


### 1.2 v0.1.1 
用户相关操作
 - [ ] 用户激活 active
    用户注册后, 向用户邮箱发送 email 验证的激活邮件, 用户从邮件中获取激活码后, 进行激活操作 
 - [ ] 用户登录与登出 login / logout 
    用户以  email + passwrd 或 用户名 + password 或 userID + password 可以进行登录, 登录成功, 给用户发送通行令牌 accessToken 
  - [ ] 用户锁定 suspend
    禁止用户业务消费行为, 一般来说, 用户在具体业务上, 会有指定的服务周期, 超出服务周期, 用户被锁定
  - [ ] 用户恢复 resume
    恢复用户授权的业务行为
  - [ ] 用户删除 deleted 
    用户删除自己或管理后台删除用户, 注意, 在商用中, 用户“删除”一般意味着用户状态修改为 deleted 并停止所有访问与所有业务, 但保留所有关联数据( 这些关联数据, 只有在一定期限后, 会手动或自动清除 purge )

## 2. 设计

从 1.1 小节, 很容易作一个简单设计

用户模型

> 1. 用户ID, 与 goim 配合, 就用 int64 吧, 一个全局唯一的正整数
> 2. email + password ,  简单点, 都是字符串,  email 格式也不验证, 其中 email 字符串长不得少于5个字符( >=5 ) , 密码不得少于6个字符( >= 6)
> 3. 用户创建时间, 简单点, 用 int64 或  UTC timestamp 
> 4. access tokdn 通告令牌, 也用字符串代替吧, 字符串长度 >=32 字符

用户操作

> 1. 用户注册  register ------> 调用 exists 检查是否重复注册, 如果不是, 创建用户ID 并保存用户信息
> 2. 检查用户是否重复存在 exists
> 3. 用户登录 login ------> 登录成功, 创建 access token
> 4. 用户登出 logout -----> 清除 access token 
> 5. 用户认证 auth -------> 调用 verify 检查 token 是否合法, 如果 token 合法则返回成功
> 6. 用户令牌验证 verify --------> 检查 token 是否存在, 并且是否相等



操作结果(状态)定义

> 1. 操作成功返回 200 状态码, 与成功信息( 文本信息)
> 2. 操作失败返回 500 状态码, 与失败原因( 文本信息)



约束与限制:

> 1. 为了快速实现原型, 用户密码直接存储, 不加密
> 2. 用户ID , 直接用用户创建的 utc 时间截( 纳秒 )

## 3.实现

用户模型与操作的 golang 实现

```
// Account  account  define
type Account struct {
	ID          int64  // 全局唯一
	Email       string // email length >=5
	Password    string // password length >=6
	CreateAt    time.Time
	AccessToken string // accessToken length >=32
}

type OperationResult struct {
	Code int    // code = 200, success;  code=500, error
	Msg  string // return operation name when success,  return error message when error
}

// AccountOperation account operation interface define
type AccountOperation interface {
	Register(email, password string) (account *Account, result OperationResult)
	Exists(email string) (exists bool, result OperationResult)
	Login(email, password string) (token string, result OperationResult)
	Logout(token string) (result OperationResult)
	Auth(token string) (pass bool, result OperationResult)
	Verify(token string) (pass bool, result OperationResult)
}
```



## 4. 测试/验证

待续...

## 5. 附注/参考

持续...
