因为把配置文件打包进了go二进制文件里，所以每次更新配置需要重新生成

先安装go-bindata

```shell script
go get -u github.com/a-urth/go-bindata/... 
```

打包配置文件为go代码
```shell script
go-bindata -pkg config .
```
