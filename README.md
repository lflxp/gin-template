# 介绍

本项目是gin框架的apiserver脚手架项目，预置内容有：

* [x] 项目模块分层：
    * [x] controller,负责处理对外接口的实现、swagger调用的代码注释
    * [x] docs，swagger项目自定义文档
    * [x] middlewares，中间件模块，纳管gin用到的中间件
        * [x] cors 跨域插件
        * [x] norouter 404插件
        * [x] prom prometheus监控插件
            * [x] 自定义prometheus metrics
        * [x] swagger 接口插件
        * [*] jwt 插件
            * [*] Authroization: Bearer $TOEKN
        * [ ] sso 单点登录插件
        * [ ] cas 权限认证插件
    * [x] model，统一模型管理文件集
    * [x] router，路由注册模块和插件注册模块
    * [x] utils，功能组件模块
        * [x] Xorm 工具
            * [x] Sqlite3
            * [x] Mysql 
            * [x] Pg
            * [x] Redis
        * [x] 缓存
            * [x] ConcurrentMap
    * [x] main.go 项目启动入口文件
* [ ] 内置前端界面管理模块
    * [ ] 前端动态页面权限
    * [ ] 基于jwt的前后端接口联调功能

本项目旨在解决项目搭建重复劳动的问题，并整合框架和常用插件，快速项目启动并避免每次新项目扣脑壳的问题。

# 使用说明

`命令模式`: 

> git clone https://github.com/lflxp/gin-template

> cd gin-template && go run main.go

`一键运行`:

> git clone https://github.com/lflxp/gin-template

> cd gin-template && make
