# teamup 微信小程序后端

## 运行


* 安装[protoc](https://github.com/protocolbuffers/protobuf)

* 安装所需要的依赖程序
```shell script
go get -u github.com/a-urth/go-bindata/...
go get -u https://github.com/rakyll/statik
go get -u https://github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
```

* 安装docker及docker-compose

* 运行项目
```shell script
make run
```

## 开源技术/框架/参考

* [micro](https://github.com/micro/micro) 微服务/GRPC框架（服务注册、发现、远程调用）
* [docker](https://www.docker.com/) 项目打包
* [docker-compose](https://github.com/docker/compose) 容器编排
* [golang/protobuf](https://github.com/golang/protobuf) 序列化
* [gorm](https://github.com/jinzhu/gorm) 一个很棒的ORM框架
* [gen](https://github.com/smallnest/gen) 生成gorm数据库model和dao代码
* [gin](https://github.com/gin-gonic/gin) 轻量的web框架（实现文档及API网关）
* [rakyll/statik](https://github.com/rakyll/statik) 把文件或目录打包进go代码（打包配置文件解决相对路径问题）
* [a-urth/go-bindata](https://github.com/a-urth/go-bindata) 将目录打包成一个FileSystem
* [mysql](https://github.com/mysql/mysql-server) 数据库
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) Standard Go Project Layout
* [swagger-ui](https://github.com/swagger-api/swagger-ui)
  
