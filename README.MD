# oldme 个人博客接口
- https://oldme.net 个人博客网站，本项目可以做为 `GoFrame` 的初学参考，会一直更新使用 `Go` 和 `GF` 的最新版本
- 使用 `Go` + `GoFrame` 开发，感谢 `GoFrame`：https://github.com/gogf/gf
- 后台登录组件使用golang-jwt/jwt：https://github.com/golang-jwt/jwt
- 前端基于 `Vue3` + `Nuxt3` 开发

# 版本信息
- Go v1.19+  
- GoFrame v2.7.1

# 功能模块
- 文章分类管理
- 文章发布
- 句子收集
- 友情链接
- 文章回复功能
- 图片文件上传
- 富文本图片文件处理
- 后台登录认证

# 特色功能设计
## 接口异常和错误的区别
业务开发中，通常有**错误**和**异常**两种 `err`，我这里暂且这么称呼，也有的称为业务异常（错误）和系统异常（异常）。错误是由用户输入不当引起的，比如说账号密码错误，这种 `err` 通常只返回给用户即可，不需要打印堆栈信息。而异常是由系统内部自发引起的，比如说 `SQL` 语句不当，这种错误需要打印堆栈信息，且不能把 `err` 返回到用户那里。`oldme-api` 合理的处理了两种不同的 `err`，他们体现在 `intrnal->pecked->err.go` 下。

## model设计
网站的前后台接口响应通常是不一样的，前台接口的响应数据总是要过滤掉一些敏感信息，比如文章的软删除字段，就不能返回给前台接口。
在 `gf` 的标准路由中，以一个结构体来定义一个接口的响应数据结构，按照简单的设计，后台和前台应各自定义一个结构体，用来处理不同的响应，但像文章标题、作者、创建时间这种信息是两者公用的，我们的两个结构体就不可避免的大量定义重复的字段。
为了解决这个问题，我们可以基于 `gf` 生成的标准数据结构，也就是 `model.entity` 下的文件。在此基础上，使用空结构体和 `omitempty` 标签来剔除掉不想返回的信息，这样就可以优雅的解决这一问题，请参考：https://github.com/oldme-git/oldme-api/blob/master/internal/model/article.go

# 部署
- git clone https://github.com/oldme-git/oldme-api
- 安装 mysql 数据库，运行 manifest/sql/oldme.sql 文件
- 安装 redis
- 修改 manifest/config下config.yaml配置文件
- 安装 go.mod 依赖
- 运行 go run main.go

# 特别注意
- Logic 和 Api 之间在应该单独定义
- Logic 弃用 `gf gen service`