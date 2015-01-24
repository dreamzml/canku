canku 基与golang的点餐系统
==========================

##使用技術
- beego
- go
- git
- mysql

## 获取安装
  执行以下命令，就能够在你的GOPATH/src/github.com/dreamzml 目录下发现canku项目
```bash
$ go get github.com/dreamzml/canku
```

## 如果初次使用beego,还需执行以下三步
- golang环境搭建
  在获取安装包之前，需搭建golang的开发环境 参考[`http://golang.in/`](golang入门-安装及环境)

- 安装git命令行工具

- 安装beego框架
```bash
$ go get github.com/astaxie/beego
```
  
## 创建数据库并同步数据

- 配置数据库链接信息（配置文件：$GOPATH/src/github.com/dreamzml/canku/conf/app.conf）
```
# MYSQL 配置
dbdriver = mysql                //数据库类型（暂时只支持mysql）
dbhost = 127.0.0.1              //数据库地址
dbname = canku                  //数据库名称（如果不存在会自动创建，如果存在会自己履盖）
dbuser = root                   //数据库用户名（建议有创建数据库及的权限）
dbpass = 123456                 //数据库密码
dbport = 3306                   //数据库端口号
```
  
- 进入项目根目录，编译项目，并执行数据库同步命令
```
$ cd $GOPATH/src/github.com/dreamzml/canku
$ go guild
$ canku -syncdb
```

## 运行项目
- 执行运行命令
```
$ cd $GOPATH/src/github.com/dreamzml/canku
$ bee run canku
```

- 浏览效果
  打开浏览器并访问 http://localhost:8080


