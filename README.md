# grpc-mysql-demo
基于protobuf的RPC实现简单mysql的增删改查操作

### 运行前准备
修改service目录下的goods_service.go 文件，将数据库的配置改成你自己的配置信息

### 运行

（1）首先执行以下操作，下载响应的包

```
go mod tidy
```

```
go mod vendor
```

（2）运行server
```
go run goods_server.go
```

（3）运行client
```
go run goods_client.go
```
