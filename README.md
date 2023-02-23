# loginsystem

//基于Go原生包"net/http"的登录注册服务//

  1.使用context保存上下文（主要是cookie）

  2.使用Linux + Docker + nginx 配置到具体域名

  3.数据库使用postgres

  4.pkg.sh打包项目
  
  5.run.sh运行项目
  
  6.自定义日志输出到Log目录下的err.log & log.log
  
  7.特点：参数校验、错误处理、*_test.go测试模块

  9.功能：/public/*是公有页面；/adm/*为私有页面，需要携带对应token（通过中间件保存在context中）
  
  10.mongodb_test目录，项目跑起来不需要，自己练习mongo语法的时候加上去的，懒得加到.gitignoe去了
