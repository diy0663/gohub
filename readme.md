学习收获

1. gin ,路由, 中间件
2. 日志包 zap
3. gorm 的简单使用 (自动迁移,数据库链接,简单的增删改查, hook )
4. cobra 命令行使用 提高效率, 配合固定的几个模板做代码生成(make 命令)
5. viper cast 搭配做配置化
6. redis 基本操作
7. jwt 包 ,图片验证码, 短信验证码
8. govalidator 请求验证器, 搭配定制规则的使用,例如 not_exists max_cn min_cn 等
9. pkg 包的引用存放
10. 仿 laravel 的代码分层: 路由->中间件->控制器->请求验证->授权策略->model->返回响应数据
11. 在启动项目的时候经过 bootstrap 里面的一系列 init 去触发配置加载
12. 定义各个全局变量,配合 sync.Once 做单例模式
13. 使用 vscode (自动填充数据结构, thunder 客户端测试 api)
14. 使用 air 热加载
