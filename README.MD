# oldme 个人博客接口
- https://oldme.net 个人博客网站，本项目可以做为 `GoFrame` 的初学参考，会一直更新使用 `Go` 和 `GF` 的最新版本
- 使用 `Go` + `GoFrame` 开发，感谢 `GoFrame`：https://github.com/gogf/gf
- 后台登录组件使用golang-jwt/jwt：https://github.com/golang-jwt/jwt
- 前端基于 `Vue3` + `Nuxt3` 开发

# 版本信息
- Go v1.22+  
- GoFrame v2.8.3+
版本信息以`go.mod`为准

# 功能模块
- 文章分类管理
- 文章发布
- 句子收集
- 友情链接
- 文章回复功能
- 图片文件上传
- 富文本图片文件处理
- 后台登录认证

# 部署
- git clone https://github.com/oldme-git/oldme-api
- 安装 mysql 数据库，运行 manifest/sql/oldme.sql 文件
- 安装 redis
- 修改 manifest/config下config.yaml配置文件
- 安装 go.mod 依赖
- 运行 go run main.go
