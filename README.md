# Developed by @[Viggo Van](mailto:wayne3van@gmail.com)

### 多仓库go语言微服务-userMicros
### github.com/wegge0857/PolyrepoGoMicros-UserMicros

### 执行命令
```bash
go get github.com/wegge0857/PolyrepoGoMicros-ApiLink
go mod tidy
```

### 添加分布式事务管理器
```bash
go get github.com/dtm-labs/client
```
### 如barrier表在当前微服务数据库
### 在对应的微服务data层
### import "github.com/dtm-labs/client/dtmcli"
### dtmcli.SetBarrierTableName("barrier")

### 接口生成 Mock 代码文件
```bash
mockery --name=UserRepo --dir=internal/biz --output=internal/biz/mocks
```
### 测试
```bash
  go test -v ./internal/biz/
```

### 运行
go run .\cmd\userMicros\ -conf .\configs\
